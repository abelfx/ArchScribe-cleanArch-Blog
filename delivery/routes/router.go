package routes

import (
	"Blog/delivery/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(blogCtrl *controllers.BlogController) (*gin.Engine) {
	r := gin.Default()
	blogRouter := r.Group("/blog")

	blogRouter.POST("/", blogCtrl.CreateBlog)
	blogRouter.PUT("/:id", blogCtrl.UpdateBlog)
	blogRouter.GET("/:id", blogCtrl.GetBlog)
	blogRouter.GET("/", blogCtrl.GetBlogs)
	blogRouter.DELETE("/:id", blogCtrl.DeleteBlog)

	return r
}