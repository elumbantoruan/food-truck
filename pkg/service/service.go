package service

import (
	"food-truck/pkg/repository"
	"food-truck/pkg/types"
)

type Service interface {
	GetNearbyLocation(latitude float64, longitude float64, radiusInMiles float64) ([]types.FoodTruckNearby, error)
	GetLocation(latitude float64, longitude float64) (*types.FoodTruckLocation, error)
	GetFoodTrucks() ([]types.FoodTruckLocation, error)
}

func NewService(repository repository.Repository) (Service, error) {
	if err := repository.BuildFoodTruckDataMap(); err != nil {
		return nil, err
	}
	return &service{
		Repository: repository,
	}, nil
}

type service struct {
	Repository repository.Repository
}

func (svc *service) GetNearbyLocation(latitude float64, longitude float64, radiusInMiles float64) ([]types.FoodTruckNearby, error) {
	return svc.Repository.GetFoodTruckNearbyLocation(latitude, longitude, radiusInMiles)
}

func (svc *service) GetLocation(latitude float64, longitude float64) (*types.FoodTruckLocation, error) {
	return svc.Repository.GetFoodTruckLocation(latitude, longitude)
}

func (svc *service) GetFoodTrucks() ([]types.FoodTruckLocation, error) {
	return svc.Repository.GetFoodTrucks()
}
