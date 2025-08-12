package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Blog represents a blog post in the system
type Blog struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Content     string             `json:"content" bson:"content"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	Tags        []string           `json:"tags" bson:"tags"`
	Likes       int                `json:"likes" bson:"likes"`
	Dislikes    int                `json:"dislikes" bson:"dislikes"`
	ViewCount   int                `json:"view_count" bson:"view_count"`
	DateCreated time.Time          `json:"date_created" bson:"date_created"`
}

// BlogInteraction tracks which user liked/disliked a blog to prevent duplicates
type BlogInteraction struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BlogID    primitive.ObjectID `bson:"blog_id" json:"blog_id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Action    string             `bson:"action" json:"action"` // "like" or "dislike"
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

// BlogRepository defines repository operations for blogs and popularity tracking
type BlogRepository interface {
	CreateBlog(blog *Blog) (*Blog, error)
	GetBlogByID(id primitive.ObjectID) (*Blog, error)
	GetAllBlogs() ([]*Blog, error)
	UpdateBlog(id primitive.ObjectID, blog *Blog) (*Blog, error)
	DeleteBlog(id primitive.ObjectID) error
	SearchBlog(title string) (*Blog, error)

	// Popularity tracking
	LikeBlog(userID, blogID primitive.ObjectID) error
	DislikeBlog(userID, blogID primitive.ObjectID) error
	IncrementViewCount(blogID primitive.ObjectID) error

	// Filtration
	FilterBlogs(tags []string, startDate, endDate *time.Time, sortBy string) ([]*Blog, error)
}

type AIService interface {
    Suggest(prompt string) (string, error)
}
