package main

import (
	"fmt"
	"sqlc-demo/internal/app"
)

func main() {
	fmt.Println("testing linter - 1")
	app.StartApp()
	// err := app.StartApp()
	// log.Println(err)
	// // if err != nil {
	// // 	log.Println(err.Error())
	// // }
}
