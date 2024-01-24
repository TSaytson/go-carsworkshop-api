package main

import (
	"fmt"
	carsController "main/controllers"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	r := httprouter.New()
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

	r.GET("/health", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "Server is running")
	})

	r.POST("/createCar", carsController.CreateCar)
	r.GET("/getCars", carsController.GetCars)

	handler := cors.Handler(r)
	port := ":3002"
	http.ListenAndServe(port, handler)
}
