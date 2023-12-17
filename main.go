package main

import (
	"log"
	"net/http"

	"github.com/gofr-dev/gofr"
	"Golang api/handlers"
)

func main() {
	// Initialize the GoFr framework
	app := gofr.New()

	// Register API handlers
	app.AddRoute("/cars", handlers.GetCars, "GET")
	app.AddRoute("/cars", handlers.CreateCar, "POST")
	app.AddRoute("/cars/{id}", handlers.GetCarByID, "GET")
	app.AddRoute("/cars/{id}", handlers.UpdateCar, "PUT")
	app.AddRoute("/cars/{id}", handlers.DeleteCar, "DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", app))
}
