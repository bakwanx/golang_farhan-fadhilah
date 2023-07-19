package main

import (
	"docker_container/route"
)

func main() {
	e := route.New()
	e.Logger.Fatal(e.Start(":8080"))
}
