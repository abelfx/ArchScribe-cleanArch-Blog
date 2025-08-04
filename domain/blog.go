package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Blog represents a blog post in the system
type Blog struct {
	ID primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	Tags []string `json:"tags" bson:"tags"`
	Likes int `json:"likes" bson:"likes"`
	Dislikes int `json:"dislikes" bson:"dislikes"`
	ViewCount int `json:"view_count" bson:"view_count"`
	DateCreated time.Time `json:"date_created" bson:"date_created"`
}

type BlogRepository interface {
	CreateBlog(blog *Blog) (*Blog, error)
	GetBlogByID(id primitive.ObjectID) (*Blog, error)
	GetAllBlogs() ([]*Blog, error)
	UpdateBlog(id primitive.ObjectID, blog *Blog) (*Blog, error)
	DeleteBlog(id primitive.ObjectID) error
	SearchBlog(id primitive.ObjectID) (*Blog, error)
	FilterBlodgsByTag(tag string) ([]*Blog, error)
	LikeBlog(id primitive.ObjectID) (*Blog, error)
	DislikeBlog(id primitive.ObjectID) (*Blog, error)
	ViewBlog(id primitive.ObjectID) (*Blog, error)
}