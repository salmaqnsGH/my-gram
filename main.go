package main

import (
	"fmt"
	database "my-gram/database"
	router "my-gram/routers"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Starting server..")
	port := os.Getenv("PORT")

	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbConf := database.Database{
		Host:      os.Getenv("DB_HOST"),
		Username:  os.Getenv("DB_USER"),
		Password:  os.Getenv("DB_PASSWORD"),
		Port:      dbPort,
		Name:      os.Getenv("DB_NAME"),
		DebugMode: os.Getenv("DEBUG_MODE"),
	}

	database.StartDB(&dbConf)

	db := database.GetDB()
	router.New(db).Run(fmt.Sprintf(":%s", port))
}
