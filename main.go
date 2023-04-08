package main

import (
	database "my-gram/database"
	router "my-gram/routers"
)

func main() {
	database.StartDB()

	db := database.GetDB()
	router.New(db).Run(":3000")
}
