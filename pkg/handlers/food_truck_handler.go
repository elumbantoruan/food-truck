package handlers

import (
	"errors"
	"food-truck/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewHandler(service service.Service) *Handler {
	hd := &Handler{
		Service: service,
	}
	return hd
}

type Handler struct {
	Service service.Service
}

func (hd *Handler) HandleGetFoodTruckNearbyLocation(c *gin.Context) {
	locationNearby, err := hd.validateLocationNearby(c)
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	foodTrucksNearbyLocation, err := hd.Service.GetNearbyLocation(locationNearby.Latitude, locationNearby.Longitude, locationNearby.RadiusInMiles)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, foodTrucksNearbyLocation)
}

func (hd *Handler) HandleGetFoodTruckLocation(c *gin.Context) {
	location, err := hd.validateLocation(c)
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	foodTrucksLocation, err := hd.Service.GetLocation(location.Latitude, location.Longitude)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, foodTrucksLocation)
}

func (hd *Handler) HandleGetFoodTrucks(c *gin.Context) {

	foodTrucks, err := hd.Service.GetFoodTrucks()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, foodTrucks)
}

func (hd *Handler) validateLocation(c *gin.Context) (*Location, error) {
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

func (hd *Handler) validateLocationNearby(c *gin.Context) (*LocationNearby, error) {
	loc, err := hd.validateLocation(c)
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
