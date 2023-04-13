package main

import (
	"fmt"
	"log"
	database "my-gram/database"
	router "my-gram/routers"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Starting server..")
	port := os.Getenv("PORT")
	database.StartDB()

	db := database.GetDB()
	router.New(db).Run(fmt.Sprintf(":%s", port))
}
