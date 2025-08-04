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
func (u *BlogUsecase) GetBlogByID(id string) (*domain.Blog, error) {
	ObjId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	return u.repo.GetBlogByID(ObjId)
}

// retrive all blogs
func (u *BlogUsecase) GetAllBlogs() ([]*domain.Blog, error){
	return u.repo.GetAllBlogs()
}

// update a blog using its id
func ( u *BlogUsecase) UpdateBlog(id string, blog *domain.Blog) (*domain.Blog, error) {
	ObjId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	return u.repo.UpdateBlog(ObjId, blog)
}

// delete a blog using its id
func ( u *BlogUsecase) DeleteBlog(id string) error {
	ObjId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}
	return u.repo.DeleteBlog(ObjId)
}