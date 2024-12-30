package main

import (
	"fmt"
	"log"
	"sqlc-demo/internal/app"
)

func main() {
	fmt.Println("testing linter")
	err := app.StartApp()
	log.Println(err)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
}
