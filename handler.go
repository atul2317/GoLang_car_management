package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"Golang api/db"
	"Golang api/models"

	"github.com/gofr-dev/gofr"
)

func GetCars(ctx *gofr.Context) {
	cars, err := db.GetAllCars()
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	ctx.JSON(http.StatusOK, cars)
}

func CreateCar(ctx *gofr.Context) {
	var newCar models.Car
	if err := ctx.BindJSON(&newCar); err != nil {
		ctx.Error(http.StatusBadRequest, "Invalid request payload")
		return
	}

	id, err := db.InsertCar(newCar)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	newCar.ID = id
	ctx.JSON(http.StatusCreated, newCar)
}

func GetCarByID(ctx *gofr.Context) {
	idStr := ctx.Params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.Error(http.StatusBadRequest, "Invalid car ID")
		return
	}

	car, err := db.GetCarByID(id)
	if err != nil {
		ctx.Error(http.StatusNotFound, "Car not found")
		return
	}

	ctx.JSON(http.StatusOK, car)
}

func UpdateCar(ctx *gofr.Context) {
	idStr := ctx.Params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.Error(http.StatusBadRequest, "Invalid car ID")
		return
	}

	var updatedCar models.Car
	if err := ctx.BindJSON(&updatedCar); err != nil {
		ctx.Error(http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = db.UpdateCar(id, updatedCar)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	updatedCar.ID = id
	ctx.JSON(http.StatusOK, updatedCar)
}

func DeleteCar(ctx *gofr.Context) {
	idStr := ctx.Params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.Error(http.StatusBadRequest, "Invalid car ID")
		return
	}

	err = db.DeleteCar(id)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	ctx.NoContent(http.StatusNoContent)
}
