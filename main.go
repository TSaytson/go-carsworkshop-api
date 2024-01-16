package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "main/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Car struct {
	Id           int
	Manufacturer string
	Model        string
	Year         int
}

func main() {
	r := mux.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	db := db.NewDatabase()

	r.HandleFunc("/createCar", func(w http.ResponseWriter, r *http.Request) {
		stmt, err := db.Prepare(`INSERT INTO cars 
		(manufacturer, model, year) values(?,?,?)`)
		checkErr(err)

		_, err = stmt.Exec("Audi", "RS Q8", 2023)
		checkErr(err)
	})

	r.HandleFunc("/getCars", func(w http.ResponseWriter, r *http.Request) {
		cars := []Car{}

		rows, err := db.Query("SELECT * FROM cars;")
		checkErr(err)

		for rows.Next() {
			car := Car{}
			rows.Scan(&car.Id, &car.Manufacturer, &car.Model, &car.Year)
			cars = append(cars, car)
		}
		rows.Close()
		res, err := json.Marshal(cars)
		fmt.Fprintf(w, string(res))
	})

	handler := cors.Handler(r)
	http.ListenAndServe(":3002", handler)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
