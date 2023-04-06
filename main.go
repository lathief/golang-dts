package main

import (
	"assinment-8/database"
	router "assinment-8/routers"
)

func main() {
	database.StartDB()

	router.New().Run(":3000")
}
