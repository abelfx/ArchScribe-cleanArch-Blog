package main

import (
	"Blog/delivery/controllers"
	"Blog/delivery/routes"
	"Blog/domain"
	"Blog/infrastructure"
	"Blog/repositories"
	"Blog/usecases"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// connect to database
	client, err := infrastructure.ConnectDB("mongodb://localhost:27017")
	if err != nil {
		log.Println("Error connecting to the database")
	}

	err = godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, relying on system environment variables")
}
	// initialize collections
	blogCollection := client.Database("Blog").Collection("blogs")
	userCollection := client.Database("User").Collection("users")

	// initialize repositories
	var blogRepo domain.BlogRepository = repositories.NewMongoBlogRepository(*blogCollection)
	var userRepo domain.UserRepository = repositories.NewMongoUserRepository(userCollection)

	// initialize AI service
	var aiService domain.AIService = infrastructure.NewMistralAIService()

	// initialize usecases
	blogUsecase := usecases.NewBlogUsecase(blogRepo, aiService)
	userUsecase := usecases.NewUserUsecase(userRepo)

	// initialize controllers
	blogController := controllers.NewTaskController(blogUsecase)
	userController := controllers.NewUserController(userUsecase)

	// setup router
	r := routes.SetUpRouter(blogController, userController, userUsecase)

	// start server
	r.Run(":3000")
}
