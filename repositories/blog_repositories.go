package repositories

import (
	"Blog/domain"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoBlogRepository struct {
	col mongo.Collection
	interactColl *mongo.Collection
}

func NewMongoBlogRepository(col mongo.Collection) *MongoBlogRepository {
	return &MongoBlogRepository{col:col, interactColl: col.Database().Collection("blog_interactions")}
}


// IncrementViewCount increases view count atomically
func (r *MongoBlogRepository) IncrementViewCount(blogID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.col.UpdateOne(ctx, bson.M{"_id": blogID}, bson.M{"$inc": bson.M{"view_count": 1}})
	return err
}

// create a task
func (r *MongoBlogRepository) CreateBlog(blog *domain.Blog) (*domain.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.col.InsertOne(ctx, blog)

	if err != nil {
		return nil, err
	}

	return blog, nil
}

// retrive a blog using its id
func (r *MongoBlogRepository) GetBlogByID(id primitive.ObjectID) (*domain.Blog, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	    defer cancel()

		var blog domain.Blog
		err := r.col.FindOne(ctx, bson.M{"_id" :id}).Decode(&blog)

		if err != nil {
			return nil, err
		}

		r.IncrementViewCount(id)

		return &blog, nil
}
	
// retrive all blogs
func (r *MongoBlogRepository) GetAllBlogs() ([]*domain.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.col.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var blogs []*domain.Blog
	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)

	}

	return blogs, nil
}

// update a blog using its id
func (r *MongoBlogRepository) UpdateBlog(id primitive.ObjectID, blog *domain.Blog) (*domain.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.col.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": blog})
	if err != nil {
		return nil, err
	}

	blog.ID = id // ensure the id is set in the returned blog
	
	return blog, nil
}

// delete a blog using its id
func (r *MongoBlogRepository) DeleteBlog(id primitive.ObjectID) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.col.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return err
	}
	return nil
}


// helper to find existing interaction
func (r *MongoBlogRepository) findInteraction(ctx context.Context, userID, blogID primitive.ObjectID) (*domain.BlogInteraction, error) {
	var bi domain.BlogInteraction
	err := r.interactColl.FindOne(ctx, bson.M{"user_id": userID, "blog_id": blogID}).Decode(&bi)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &bi, nil
}

// LikeBlog implements logic to prevent duplicate likes and handle switch from dislike to like
func (r *MongoBlogRepository) LikeBlog(userID, blogID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// check if blog exists
	var blog domain.Blog
	if err := r.col.FindOne(ctx, bson.M{"_id": blogID}).Decode(&blog); err != nil {
		return errors.New("blog not found")
	}

	// check previous interaction
	prev, err := r.findInteraction(ctx, userID, blogID)
	if err != nil {
		return err
	}

	if prev != nil && prev.Action == "like" {
		return errors.New("user already liked this blog")
	}

	// start a session-less update sequence
	if prev != nil && prev.Action == "dislike" {
		// remove dislike record and decrement dislike counter
		_, err := r.interactColl.DeleteOne(ctx, bson.M{"_id": prev.ID})
		if err != nil {
			return err
		}
		_, err = r.col.UpdateOne(ctx, bson.M{"_id": blogID}, bson.M{"$inc": bson.M{"dislikes": -1}})
		if err != nil {
			return err
		}
	}

	// insert like interaction
	bi := domain.BlogInteraction{
		ID:        primitive.NewObjectID(),
		BlogID:    blogID,
		UserID:    userID,
		Action:    "like",
		CreatedAt: time.Now(),
	}
	if _, err := r.interactColl.InsertOne(ctx, bi); err != nil {
		return err
	}

	r.IncrementViewCount(blogID)

	// increment likes on blog
	_, err = r.col.UpdateOne(ctx, bson.M{"_id": blogID}, bson.M{"$inc": bson.M{"likes": 1}})
	return err
}

// DislikeBlog similar to LikeBlog
func (r *MongoBlogRepository) DislikeBlog(userID, blogID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var blog domain.Blog
	if err := r.col.FindOne(ctx, bson.M{"_id": blogID}).Decode(&blog); err != nil {
		return errors.New("blog not found")
	}

	prev, err := r.findInteraction(ctx, userID, blogID)
	if err != nil {
		return err
	}

	if prev != nil && prev.Action == "dislike" {
		return errors.New("user already disliked this blog")
	}

	if prev != nil && prev.Action == "like" {
		_, err := r.interactColl.DeleteOne(ctx, bson.M{"_id": prev.ID})
		if err != nil {
			return err
		}
		_, err = r.col.UpdateOne(ctx, bson.M{"_id": blogID}, bson.M{"$inc": bson.M{"likes": -1}})
		if err != nil {
			return err
		}
	}

	bi := domain.BlogInteraction{
		ID:        primitive.NewObjectID(),
		BlogID:    blogID,
		UserID:    userID,
		Action:    "dislike",
		CreatedAt: time.Now(),
	}
	if _, err := r.interactColl.InsertOne(ctx, bi); err != nil {
		return err
	}

	r.IncrementViewCount(blogID)

	_, err = r.col.UpdateOne(ctx, bson.M{"_id": blogID}, bson.M{"$inc": bson.M{"dislikes": 1}})
	return err
}



// FilterBlogs supports tags, optional date range, and sorting by date or popularity
func (r *MongoBlogRepository) FilterBlogs(tags []string, startDate, endDate *time.Time, sortBy string) ([]*domain.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{}
	if len(tags) > 0 {
		filter["tags"] = bson.M{"$all": tags}
	}

	if startDate != nil || endDate != nil {
		rangeQuery := bson.M{}
		if startDate != nil {
			rangeQuery["$gte"] = *startDate
		}
		if endDate != nil {
			rangeQuery["$lte"] = *endDate
		}
		filter["date_created"] = rangeQuery
	}

	findOptions := options.Find()
	// sort
	switch sortBy {
	case "popularity":
		// popularity: likes + view_count (descending)
		findOptions.SetSort(bson.D{{"likes", -1}, {"view_count", -1}, {"date_created", -1}})
	case "date":
		findOptions.SetSort(bson.D{{"date_created", -1}})
	default:
		findOptions.SetSort(bson.D{{"date_created", -1}})
	}

	cursor, err := r.col.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var blogs []*domain.Blog
	for cursor.Next(ctx) {
		var b domain.Blog
		if err := cursor.Decode(&b); err != nil {
			return nil, err
		}
		blogs = append(blogs, &b)
	}

	return blogs, nil
}

func (r *MongoBlogRepository) SearchBlog(title string) (*domain.Blog, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    filter := bson.M{
        "title": bson.M{
            "$regex":   title,
            "$options": "i", // case-insensitive
        },
    }

    var blog domain.Blog
    err := r.col.FindOne(ctx, filter).Decode(&blog)
    if err != nil {
        return nil, err
    }
    return &blog, nil
}
