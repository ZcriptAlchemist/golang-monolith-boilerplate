package main

import (
	"fmt"
	"log"
	"sqlc-demo/internal/app"
)

func main() {
	fmt.Println("okay okay")
	err := app.StartApp()
	if err != nil {
		log.Println(err.Error())
	}
}
