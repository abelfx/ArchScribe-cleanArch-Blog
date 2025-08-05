package controllers

import (
	"Blog/domain"
	"Blog/usecases"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogController struct {
	blogUsecase *usecases.BlogUsecase
}

func NewTaskController(u *usecases.BlogUsecase) *BlogController {
	return &BlogController{blogUsecase: u}
}

// create a blog
func (ctrl *BlogController) CreateBlog(c *gin.Context) {
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, ok := userIDValue.(string)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}

	userId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}


	var input struct {
		Title   string             `json:"title"`
		Content string             `json:"content"`
		Tags    []string           `json:"tags"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	blog := domain.Blog{
		ID:          primitive.NewObjectID(), 
		Title:       input.Title,
		Content:     input.Content,
		UserID:      userId,
		Tags:        input.Tags,
		Likes:       0,                      
		Dislikes:    0,                      
		ViewCount:   0,                      
		DateCreated: time.Now(),            
	}


	CreatedBlog, err := ctrl.blogUsecase.CreateBlog(&blog)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	    return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "blog created successfully", "Created_Blog": CreatedBlog})

}

// get a blog by id
func(ctrl *BlogController) GetBlog(c *gin.Context) {
	id := c.Param("id")
	blog, err := ctrl.blogUsecase.GetBlogByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "blog not found"})
        return
	}

	if blog == nil {
			c.JSON(http.StatusOK, gin.H{"message": "Blog not found"})
	}
	c.JSON(http.StatusOK, blog)
} 

// get all blogs
func(ctrl *BlogController) GetBlogs(c *gin.Context) {
	blogs, err := ctrl.blogUsecase.GetAllBlogs()

	if err != nil {
	    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(blogs) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No blogs posted"})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// update a blog
func (ctrl *BlogController) UpdateBlog(c *gin.Context) {
	id := c.Param("id")
	var blog domain.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UpdatedBlog, err := ctrl.blogUsecase.UpdateBlog(id, &blog)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	    return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "Blog updated successfully",
		"updated_blog": UpdatedBlog,
	})

}

// delete blog
func(ctrl *BlogController) DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.blogUsecase.DeleteBlog(id)
	
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
        return
	}
	c.JSON(http.StatusOK, gin.H{"message": "blog deleted successfully"})
} 
