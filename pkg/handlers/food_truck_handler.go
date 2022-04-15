package handlers

import (
	"errors"
	"food-truck/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewFoodTruckHandler(foodTruckService service.FoodTruckService) *FoodTruckHandler {
	ft := &FoodTruckHandler{
		FoodTruckService: foodTruckService,
	}
	return ft
}

type FoodTruckHandler struct {
	FoodTruckService service.FoodTruckService
}

func (ft *FoodTruckHandler) HandleGetFoodTruckNearbyLocation(c *gin.Context) {
	locationNearby, err := ft.validateLocationNearby(c)
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	foodTrucksNearbyLocation, err := ft.FoodTruckService.GetFoodTruckNearbyLocation(locationNearby.Latitude, locationNearby.Longitude, locationNearby.RadiusInMiles)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, foodTrucksNearbyLocation)
}

func (ft *FoodTruckHandler) HandleGetFoodTruckLocation(c *gin.Context) {
	location, err := ft.validateLocation(c)
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	foodTrucksLocation, err := ft.FoodTruckService.GetFoodTruckLocation(location.Latitude, location.Longitude)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, foodTrucksLocation)
}

func (ft *FoodTruckHandler) HandleGetFoodTrucks(c *gin.Context) {

	foodTrucks, err := ft.FoodTruckService.GetFoodTrucks()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, foodTrucks)
}

func (ft *FoodTruckHandler) validateLocation(c *gin.Context) (*Location, error) {
	lat, ok := c.Params.Get("latitude")
	if !ok {
		return nil, errors.New("latitude param is empty")
	}
	latitude, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return nil, errors.New("latitude value is not valid")
	}
	long, ok := c.Params.Get("longitude")
	if !ok {
		return nil, errors.New("longitude param is empty")
	}
	longitude, err := strconv.ParseFloat(long, 64)
	if err != nil {
		return nil, errors.New("longitude value is not valid")
	}
	return &Location{
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}

func (ft *FoodTruckHandler) validateLocationNearby(c *gin.Context) (*LocationNearby, error) {
	loc, err := ft.validateLocation(c)
	if err != nil {
		return nil, err
	}
	rad, ok := c.Params.Get("radius")
	if !ok {
		return nil, errors.New("rad param is empty")
	}
	radius, err := strconv.ParseFloat(rad, 64)
	if err != nil {
		return nil, errors.New("radius value is not valid")
	}
	return &LocationNearby{
		Location:      *loc,
		RadiusInMiles: radius,
	}, nil
}

type Location struct {
	Latitude  float64
	Longitude float64
}

type LocationNearby struct {
	Location
	RadiusInMiles float64
}
