package main

import (
	"Blog/delivery/controllers"
	"Blog/delivery/routes"
	"Blog/domain"
	"Blog/infrastructure"
	"Blog/repositories"
	"Blog/usecases"
	"log"
)

func main() {
	// connect to database
	client, err := infrastructure.ConnectDB("mongodb://localhost:27017")
	if err != nil {
		log.Println("Error connecting to the database")
	}

	// initialize collection
	blogCollection := client.Database("Blog").Collection("blogs")
	userCollection := client.Database("User").Collection("users")

	// initialize Repository
	var blogRepo domain.BlogRepository = repositories.NewMongoBlogRepository(*blogCollection)
	var userRepo domain.UserRepository = repositories.NewMongoUserRepository(userCollection)

	// initialzie Usecase
	blogUsecase := usecases.NewBlogUsecase(blogRepo)
	userUsecase := usecases.NewUserUsecase(userRepo)

	// initialize Controller
	blogController := controllers.NewTaskController(blogUsecase)
	userController := controllers.NewUserController(userUsecase)

	r := routes.SetUpRouter(blogController, userController, userUsecase)

	// start server
	r.Run(":3000")

}