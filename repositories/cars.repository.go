package repositories

import (
	"main/database"
	"main/models"
	"main/utils"
)

var db = database.NewDatabase()

func GetCars(cars []models.Car) {
	rows, err := db.Query("SELECT * FROM cars;")
	utils.CheckErr(err)

	for rows.Next() {
		car := models.Car{}
		rows.Scan(&car.Id, &car.Manufacturer, &car.Model, &car.Year)
		cars = append(cars, car)
	}
	rows.Close()
}
