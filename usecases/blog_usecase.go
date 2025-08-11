package usecases

import (
	"Blog/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// struct to implement usecase
type BlogUsecase struct {
	repo domain.BlogRepository
	aiService domain.AIService
}

// constructor for BlogUsecase
func NewBlogUsecase(repo domain.BlogRepository,  aiService domain.AIService) *BlogUsecase {
	return &BlogUsecase{repo: repo, aiService: aiService}
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


// Popularity actions
func (u *BlogUsecase) LikeBlog(userID, blogID string) error {
	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	bid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}
	return u.repo.LikeBlog(uid, bid)
}

func (u *BlogUsecase) DislikeBlog(userID, blogID string) error {
	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	bid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}
	return u.repo.DislikeBlog(uid, bid)
}

// FilterBlogs accepts optional dates (send nil to ignore)
func (u *BlogUsecase) FilterBlogs(tags []string, start, end *time.Time, sortBy string) ([]*domain.Blog, error) {
	return u.repo.FilterBlogs(tags, start, end, sortBy)
}

func (u *BlogUsecase) SuggestContent(prompt string) (string, error) {
    return u.aiService.Suggest(prompt)
}