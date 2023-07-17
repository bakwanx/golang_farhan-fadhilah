package main

import (
	"orm_and_mvc/config"
	"orm_and_mvc/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8080")
}
