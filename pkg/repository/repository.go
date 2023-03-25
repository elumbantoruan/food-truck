package repository

import (
	"encoding/csv"
	"errors"
	"fmt"
	"food-truck/pkg/types"
	"net/http"
	"strconv"

	"github.com/umahmood/haversine"
)

type Repository interface {
	BuildFoodTruckDataMap() error
	GetFoodTruckNearbyLocation(lat float64, lon float64, radiusMiles float64) ([]types.FoodTruckNearby, error)
	GetFoodTruckLocation(latitude float64, longitude float64) (*types.FoodTruckLocation, error)
	GetFoodTrucks() ([]types.FoodTruckLocation, error)
}

func NewRepository(url string) Repository {
	return &Storage{
		MemoryStore: make(map[string]types.FoodTruckLocation),
		url:         url,
	}
}

type Storage struct {
	MemoryStore map[string]types.FoodTruckLocation
	url         string
}

func genKey(lat float64, long float64) string {
	return fmt.Sprintf("(%v,%v)", lat, long)
}

func (st *Storage) BuildFoodTruckDataMap() error {
	resp, err := http.Get(st.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var foodTrucksMap = make(map[string]types.FoodTruckLocation)

	csvReader := csv.NewReader(resp.Body)
	var firstLine = 0 // header
	for {
		data, err := csvReader.Read()
		if err != nil {
			break
		}
		if firstLine == 0 {
			firstLine++
			continue
		}
		foodTruck := parseFoodTruckLocation(data)
		key := genKey(foodTruck.Latitude, foodTruck.Longitude)
		foodTrucksMap[key] = foodTruck
	}
	st.MemoryStore = foodTrucksMap
	return nil

}

// GetFoodTruckNearbyLocation given the current location of latitude, longitude, and radius.
// It utilizes haversine formula https://en.wikipedia.org/wiki/Haversine_formula
// to get the list of food trucks given the radius of current lat/lon (location)
// This is the poor implementation as it scans the whole records (O n)
// Obviously there is a better way, but the purpose of this is just to provide a good user experience
func (st *Storage) GetFoodTruckNearbyLocation(lat float64, lon float64, radiusMiles float64) ([]types.FoodTruckNearby, error) {
	var foodTrucksNearbies []types.FoodTruckNearby
	origin := haversine.Coord{Lat: lat, Lon: lon}
	for _, foodtruck := range st.MemoryStore {
		if foodtruck.Latitude == float64(0) || foodtruck.Longitude == float64(0) {
			continue
		}
		target := haversine.Coord{Lat: foodtruck.Latitude, Lon: foodtruck.Longitude}
		miles, _ := haversine.Distance(origin, target)
		if miles <= radiusMiles {
			foodTruckNearby := types.FoodTruckNearby{
				FoodTruckLocation: foodtruck,
				DistanceInMiles:   miles,
			}
			foodTrucksNearbies = append(foodTrucksNearbies, foodTruckNearby)
		}
	}
	return foodTrucksNearbies, nil
}

// GetFoodTruckLocation given latitude and longitude
// Time complexity is 0(1) since data is stored in map
func (st *Storage) GetFoodTruckLocation(latitude float64, longitude float64) (*types.FoodTruckLocation, error) {
	key := genKey(latitude, longitude)
	if fd, ok := st.MemoryStore[key]; ok {
		return &fd, nil
	}
	return nil, errors.New("not found")
}

func (st *Storage) GetFoodTrucks() ([]types.FoodTruckLocation, error) {
	var foodTrucksLocation []types.FoodTruckLocation
	for _, v := range st.MemoryStore {
		foodTrucksLocation = append(foodTrucksLocation, v)
	}
	return foodTrucksLocation, nil
}

func parseFoodTruckLocation(items []string) types.FoodTruckLocation {
	lat, _ := strconv.ParseFloat(items[14], 64)
	long, _ := strconv.ParseFloat(items[15], 64)

	return types.FoodTruckLocation{
		FoodTruckName:       items[1],
		FacilityType:        items[2],
		LocationDescription: items[4],
		Address:             items[5],
		FoodItems:           items[11],
		Latitude:            lat,
		Longitude:           long,
		Schedules:           items[17],
	}
}
