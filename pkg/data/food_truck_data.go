package data

type FoodTruckLocation struct {
	FoodTruckName       string  `json:"food_truck_name"`
	FacilityType        string  `json:"facility_type"`
	LocationDescription string  `json:"location_description"`
	Address             string  `json:"address"`
	FoodItems           string  `json:"food_items"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	Schedules           string  `json:"schedules"`
}

type FoodTruckNearby struct {
	FoodTruckLocation
	DistanceInMiles float64 `json:"distance_in_miles"`
}
