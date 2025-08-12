package routes

import (
	"Blog/delivery/controllers"
	"Blog/infrastructure"
	"Blog/usecases"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(blogCtrl *controllers.BlogController, userCtrl *controllers.UserController, useCase *usecases.UserUsecase) (*gin.Engine) {
	r := gin.Default()

	// register and login (public routes)
	authRouter := r.Group("/auth")
	authRouter.POST("/signup", userCtrl.Signup)
	authRouter.POST("/login", userCtrl.Login)
	authRouter.POST("/forgot-password", userCtrl.ForgotPassword)
	authRouter.POST("/reset-password", userCtrl.ResetPassword)
	authRouter.POST("/logout", infrastructure.AuthMiddleware(), userCtrl.Logout)

	// admin routers (required authentication and admin role)
	adminRouter := r.Group("/admin")
	adminRouter.Use(infrastructure.AuthMiddleware())
	// adminRouter.Use(infrastructure.AdminOnly(useCase))
	adminRouter.GET("/users", userCtrl.GetUsers)
	adminRouter.POST("/users/promote", userCtrl.PromoteUser)
	adminRouter.DELETE("/users/:id", userCtrl.DeleteUser)


	blogRouter := r.Group("/blog")
	blogRouter.POST("/", infrastructure.AuthMiddleware(), blogCtrl.CreateBlog)
	blogRouter.PUT("/:id", infrastructure.AuthMiddleware(), blogCtrl.UpdateBlog)
	blogRouter.GET("/:id", blogCtrl.GetBlog)
	blogRouter.GET("/", blogCtrl.GetBlogs)
	blogRouter.POST("/blogs/search", blogCtrl.SearchBlog)
	blogRouter.DELETE("/:id", infrastructure.AuthMiddleware(), blogCtrl.DeleteBlog)
	blogRouter.POST("/:id/like",infrastructure.AuthMiddleware(), blogCtrl.LikeBlog)
	blogRouter.POST("/:id/dislike", infrastructure.AuthMiddleware(),blogCtrl.DislikeBlog)
	blogRouter.POST("/filter", blogCtrl.FilterBlogs)
	blogRouter.POST("/suggest", infrastructure.AuthMiddleware(), blogCtrl.SuggestBlog)

	return r
}