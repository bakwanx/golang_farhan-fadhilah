package main

import (
	"clean_architecture/route"
)

func main() {

	app := route.New()
	app.Start(":8080")

}
