package main

import (
	"log"

	"github.com/glopezep/arithmetic-calculator/internal/application"
	"github.com/glopezep/arithmetic-calculator/internal/interfaces/lambda/handlers"
)

var app *application.Application

func init() {
	a, err := application.NewApplication()
	if err != nil {
		log.Fatal(err)
	}

	app = a
}

func main() {
	handlers.StartListRecordsHandler(app)
}
