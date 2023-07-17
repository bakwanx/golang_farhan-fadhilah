package main

import (
	"clean_architecture/config"
	"clean_architecture/route"
)

func main() {
	config.ConnectDB()
	app := route.New()
	app.Start(":8080")

}
