package usecases

import (
	"Blog/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// struct to implement usecase
type TaskUsecase struct {
	repo domain.BlogRepository
}

// constructor for TaskUsecase
func NewTaskUsecase(repo domain.BlogRepository) *TaskUsecase {
	return &TaskUsecase{repo: repo}
}


// create blog
func(u *TaskUsecase) CreateBlog(blog *domain.Blog) (*domain.Blog, error) {
	return u.repo.CreateBlog(blog)
}

// retrive a blog using its id
func (u *TaskUsecase) GetBlogByID(id primitive.ObjectID) (*domain.Blog, error) {
	return u.repo.GetBlogByID(id)
}

// retrive all blogs
func (u *TaskUsecase) GetAllBlogs() ([]*domain.Blog, error){
	return u.repo.GetAllBlogs()
}

// update a blog using its id
func ( u *TaskUsecase) UpdateBlog(id primitive.ObjectID, blog *domain.Blog) (*domain.Blog, error) {
	return u.repo.UpdateBlog(id, blog)
}

// delete a blog using its id
func ( u *TaskUsecase) DeleteBlog(id primitive.ObjectID) error {
	return u.repo.DeleteBlog(id)
}