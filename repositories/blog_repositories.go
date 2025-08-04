package repositories

import (
	"Blog/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoBlogRepository struct {
	col mongo.Collection
}

func NewBlogUsecase(col mongo.Collection) *MongoBlogRepository {
	return &MongoBlogRepository{col:col}
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
