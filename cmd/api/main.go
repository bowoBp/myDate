package main

import (
	"github.com/bowoBp/myDate/pkg/api"
	"log"
)

func main() {
	app := api.Default()
	err := app.Start()
	if err != nil {
		log.Print(err)
		panic(err)
	}
}
