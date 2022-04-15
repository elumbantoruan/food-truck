package service

import (
	"food-truck/pkg/data"
	"food-truck/pkg/repository"
)

type FoodTruckService interface {
	GetFoodTruckNearbyLocation(latitude float64, longitude float64, radiusInMiles float64) ([]data.FoodTruckNearby, error)
	GetFoodTruckLocation(latitude float64, longitude float64) (*data.FoodTruckLocation, error)
	GetFoodTrucks() ([]data.FoodTruckLocation, error)
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
	return ft.Repository.GetFoodTruckNearbyLocation(latitude, longitude, radiusInMiles)
}

func (ft *FoodTruckProvider) GetFoodTruckLocation(latitude float64, longitude float64) (*data.FoodTruckLocation, error) {
	return ft.Repository.GetFoodTruckLocation(latitude, longitude)
}

func (ft *FoodTruckProvider) GetFoodTrucks() ([]data.FoodTruckLocation, error) {
	return ft.Repository.GetFoodTrucks()
}
