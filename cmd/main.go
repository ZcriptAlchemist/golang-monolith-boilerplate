package main

import (
	"log"
	"sqlc-demo/internal/app"
)

func main() {
	err := app.StartApp()
	if err != nil {
		log.Println(err.Error())
	}
}
