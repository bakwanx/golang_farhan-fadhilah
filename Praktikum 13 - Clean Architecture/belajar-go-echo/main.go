package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/route"
)

func main() {
	config.ConnectDB()
	app := route.New()
	app.Start(":8080")

}
