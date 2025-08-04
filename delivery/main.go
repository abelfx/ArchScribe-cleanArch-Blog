package main

import (
	"Blog/delivery/controllers"
	"Blog/delivery/routes"
	"Blog/domain"
	"Blog/infrastructrue"
	"Blog/repositories"
	"Blog/usecases"
	"log"
)

func main() {
	// connect to database
	client, err := infrastructrue.ConnectDB("mongodb://localhost:27017")
	if err != nil {
		log.Println("Error connecting to the database")
	}

	// initialize collection
	blogCollection := client.Database("Blog").Collection("blogs")

	// initialize Repository
	var blogRepo domain.BlogRepository = repositories.NewMongoBlogRepository(*blogCollection)

	// initialzie Usecase
	blogUsecase := usecases.NewBlogUsecase(blogRepo)

	// initialize Controller
	blogController := controllers.NewTaskController(blogUsecase)

	r := routes.SetUpRouter(blogController)

	// start server
	r.Run(":3000")

}