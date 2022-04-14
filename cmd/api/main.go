package main

import (
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

	// nearbyTrucks, err := srv.GetFoodTruckNearbyLocation(float64(37.809639), float64(-122.410248), float64(1))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, nearbyTruck := range nearbyTrucks {
	// 	fmt.Println(nearbyTruck)
	// }

	// location, err := srv.GetFoodTruckLocation(float64(37.724297778527635), float64(-122.45937730954839))
	// fmt.Println(location)

	handler := handlers.NewFoodTruckHandler(srv)

	route := gin.New()
	route.GET("/:latitude/:longitude", handler.HandleGetFoodTruckLocation)
	route.GET("/:latitude/:longitude/:radius", handler.HandleGetFoodTruckNearbyLocation)
	route.Run(":8088")

}
