package controllers

import (
	"encoding/json"
	"fmt"
	"main/database"
	"main/models"
	"main/repositories"
	"main/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateCar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	stmt, err := database.NewDatabase().Prepare(`INSERT INTO cars
	(manufacturer, model, year) values(?,?,?)`)
	utils.CheckErr(err)

	_, err = stmt.Exec("Audi", "RS Q8", 2023)
	utils.CheckErr(err)
}

func GetCars(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cars := []models.Car{}

	repositories.GetCars(cars)
	res, err := json.Marshal(cars)
	utils.CheckErr(err)
	fmt.Fprintf(w, string(res))
}
