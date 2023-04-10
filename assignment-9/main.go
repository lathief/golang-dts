package main

import (
	"assignment-9/database"
	router "assignment-9/routers"
)

func main() {
	database.StartDB()

	router.New().Run(":3000")
}
