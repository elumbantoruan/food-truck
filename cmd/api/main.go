package main

import (
	"fmt"
	"food-truck/pkg/handlers"
	"food-truck/pkg/repository"
	"food-truck/pkg/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	url := "https://data.sfgov.org/api/views/rqzj-sfat/rows.csv"

	repo := repository.NewFoodTruckRepository(url)

	srv, err := service.NewFoodTruckService(repo)
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.NewFoodTruckHandler(srv)

	var text = `Welcome to Food Truck Service
	The service opens up port 8088 and provides the following endpoints:
	GET http://localhost:8088/foodtrucks/:latitude/:longitude/:radius (for Proximity)
	GET http://localhost:8088/foodtrucks/:latitude/:longitude (for exact location)
	GET http://localhost:8088/foodtrucks/ (for list of FoodTrucks)
	`

	fmt.Println(text)

	gin.SetMode(gin.ReleaseMode)
	route := gin.New()
	route.GET("/foodtrucks/:latitude/:longitude", handler.HandleGetFoodTruckLocation)
	route.GET("/foodtrucks/:latitude/:longitude/:radius", handler.HandleGetFoodTruckNearbyLocation)
	route.GET("/foodtrucks/", handler.HandleGetFoodTrucks)
	route.Run(":8088")

}
