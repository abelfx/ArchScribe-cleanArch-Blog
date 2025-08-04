package usecases

import (
	"Blog/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// struct to implement usecase
type BlogUsecase struct {
	repo domain.BlogRepository
}

// constructor for BlogUsecase
func NewBlogUsecase(repo domain.BlogRepository) *BlogUsecase {
	return &BlogUsecase{repo: repo}
}


// create blog
func(u *BlogUsecase) CreateBlog(blog *domain.Blog) (*domain.Blog, error) {
	return u.repo.CreateBlog(blog)
}

// retrive a blog using its id
func (u *BlogUsecase) GetBlogByID(id primitive.ObjectID) (*domain.Blog, error) {
	return u.repo.GetBlogByID(id)
}

// retrive all blogs
func (u *BlogUsecase) GetAllBlogs() ([]*domain.Blog, error){
	return u.repo.GetAllBlogs()
}

// update a blog using its id
func ( u *BlogUsecase) UpdateBlog(id primitive.ObjectID, blog *domain.Blog) (*domain.Blog, error) {
	return u.repo.UpdateBlog(id, blog)
}

// delete a blog using its id
func ( u *BlogUsecase) DeleteBlog(id primitive.ObjectID) error {
	return u.repo.DeleteBlog(id)
}