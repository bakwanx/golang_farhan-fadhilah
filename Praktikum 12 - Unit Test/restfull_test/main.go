package main

import (
	"restfull_test/config"
	"restfull_test/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8080")
}
