package service

import (
	"errors"
	"food-truck/pkg/data"
	"food-truck/pkg/repository"
)

type FoodTruckService interface {
	GetFoodTruckNearbyLocation(latitude float64, longitude float64, radiusInMiles float64) ([]data.FoodTruckNearby, error)
	GetFoodTruckLocation(latitude float64, longitude float64) (*data.FoodTruckLocation, error)
}

func NewFoodTruckService(repository repository.FoodTruckRepository) (FoodTruckService, error) {
	if err := repository.BuildFoodTruckDataMap(); err != nil {
		return nil, err
	}
	return &FoodTruckProvider{
		Repository: repository,
	}, nil
}

type FoodTruckProvider struct {
	Repository repository.FoodTruckRepository
}

func (ft *FoodTruckProvider) GetFoodTruckNearbyLocation(latitude float64, longitude float64, radiusInMiles float64) ([]data.FoodTruckNearby, error) {
	// service validates latitude and longitude since we're not serving Food Truck in Atlantic ocean :)
	if latitude == float64(0) {
		return nil, errors.New("invalid latitude")
	}
	if longitude == float64(0) {
		return nil, errors.New("invalid longitude")
	}
	return ft.Repository.GetFoodTruckNearbyLocation(latitude, longitude, radiusInMiles), nil
}

func (ft *FoodTruckProvider) GetFoodTruckLocation(latitude float64, longitude float64) (*data.FoodTruckLocation, error) {
	// service validates latitude and longitude since we're not serving Food Truck in Atlantic ocean :)
	if latitude == float64(0) {
		return nil, errors.New("invalid latitude")
	}
	if longitude == float64(0) {
		return nil, errors.New("invalid longitude")
	}
	return ft.Repository.GetFoodTruckLocation(latitude, longitude)
}
