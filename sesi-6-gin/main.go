package main

import (
	"sesi-6-gin/routers"
)

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
