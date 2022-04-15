package handlers

import (
	"encoding/json"
	"food-truck/pkg/data"
	"food-truck/pkg/repository"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFoodTruckHandler_HandleGetFoodTruckNearbyLocation(t *testing.T) {

	var repo MockRepository
	var srv = MockService{Repository: &repo}
	var ft = NewFoodTruckHandler(&srv)

	r := getRouter()

	r.GET("/foodtrucks/:latitude/:longitude/:radius", ft.HandleGetFoodTruckNearbyLocation)
	req, _ := http.NewRequest("GET", "/foodtrucks/1.05/2.05/1", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		assert.NoError(t, err)
		var foodTrucksNearby []data.FoodTruckNearby
		json.Unmarshal(p, &foodTrucksNearby)

		expectedData, _ := repo.GetFoodTruckNearbyLocation(1.05, 5.05, 1)
		assert.Equal(t, 3, len(foodTrucksNearby))
		assert.EqualValues(t, expectedData, foodTrucksNearby)
		assert.Equal(t, true, statusOK)
		return statusOK
	})
}

func TestFoodTruckHandler_HandleGetFoodTruckNearbyLocation_InvalidRadius(t *testing.T) {

	var repo MockRepository
	var srv = MockService{Repository: &repo}
	var ft = NewFoodTruckHandler(&srv)

	r := getRouter()

	r.GET("/foodtrucks/:latitude/:longitude/:radius", ft.HandleGetFoodTruckNearbyLocation)
	req, _ := http.NewRequest("GET", "/foodtrucks/1.05/2.05/invalidradius", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusBadRequest := w.Code == http.StatusBadRequest

		assert.Equal(t, true, statusBadRequest)
		return statusBadRequest
	})
}

func TestFoodTruckHandler_HandleGetFoodTruckLocation(t *testing.T) {

	var repo MockRepository
	var srv = MockService{Repository: &repo}
	var ft = NewFoodTruckHandler(&srv)

	r := getRouter()

	r.GET("/foodtrucks/:latitude/:longitude", ft.HandleGetFoodTruckLocation)
	req, _ := http.NewRequest("GET", "/foodtrucks/1.05/2.05", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		assert.NoError(t, err)
		var foodTruck data.FoodTruckLocation
		json.Unmarshal(p, &foodTruck)

		expectedData, _ := repo.GetFoodTruckLocation(1.05, 5.05)
		assert.NotNil(t, foodTruck)
		assert.EqualValues(t, expectedData, &foodTruck)
		assert.Equal(t, true, statusOK)
		return statusOK
	})
}

func TestFoodTruckHandler_HandleGetFoodTrucksLocation(t *testing.T) {

	var repo MockRepository
	var srv = MockService{Repository: &repo}
	var ft = NewFoodTruckHandler(&srv)

	r := getRouter()

	r.GET("/foodtrucks/:latitude/:longitude", ft.HandleGetFoodTrucks)
	req, _ := http.NewRequest("GET", "/foodtrucks/1.05/2.05", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		assert.NoError(t, err)
		var foodTruck []data.FoodTruckLocation
		json.Unmarshal(p, &foodTruck)

		expectedData, _ := repo.GetFoodTrucks()
		assert.NotNil(t, foodTruck)
		assert.EqualValues(t, expectedData, foodTruck)
		assert.Equal(t, true, statusOK)
		return statusOK
	})
}

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func getRouter() *gin.Engine {
	return gin.Default()
}

func createMockFoodTrucksNearby() []data.FoodTruckNearby {
	text := `[
		{
			"food_truck_name": "Datam SF LLC dba Anzu To You",
			"facility_type": "Truck",
			"location_description": "TAYLOR ST: BAY ST to NORTH POINT ST (2500 - 2599)",
			"address": "2535 TAYLOR ST",
			"food_items": "Asian Fusion - Japanese Sandwiches/Sliders/Misubi",
			"latitude": 37.805885350100986,
			"longitude": -122.41594524663745,
			"schedules": "",
			"distance_in_miles": 0.4048767783754979
		},
		{
			"food_truck_name": "Philz Coffee Truck",
			"facility_type": "Truck",
			"location_description": "MONTGOMERY ST: COLUMBUS AVE \\ WASHINGTON ST to JACKSON ST (700 - 799)",
			"address": "735 MONTGOMERY ST",
			"food_items": "Hot coffee: iced coffee: hot chocolate: tea: pastries",
			"latitude": 37.79621549659414,
			"longitude": -122.40375455824538,
			"schedules": "",
			"distance_in_miles": 0.9927219402886978
		},
		{
			"food_truck_name": "Wu Wei LLC dba MoBowl",
			"facility_type": "Truck",
			"location_description": "FRONT ST: VALLEJO ST to GREEN ST (900 - 999)",
			"address": "900 FRONT ST",
			"food_items": "Various types of meat: veggie: and seafood bowls.",
			"latitude": 37.80022091559222,
			"longitude": -122.39956901857713,
			"schedules": "",
			"distance_in_miles": 0.8735059296153723
		}
	]`
	var records []data.FoodTruckNearby
	_ = json.Unmarshal([]byte(text), &records)
	return records
}

func createMockFoodTrucksLocation() []data.FoodTruckLocation {
	text := `[
		{
			"food_truck_name": "Datam SF LLC dba Anzu To You",
			"facility_type": "Truck",
			"location_description": "TAYLOR ST: BAY ST to NORTH POINT ST (2500 - 2599)",
			"address": "2535 TAYLOR ST",
			"food_items": "Asian Fusion - Japanese Sandwiches/Sliders/Misubi",
			"latitude": 37.805885350100986,
			"longitude": -122.41594524663745,
			"schedules": ""
		},
		{
			"food_truck_name": "Philz Coffee Truck",
			"facility_type": "Truck",
			"location_description": "MONTGOMERY ST: COLUMBUS AVE \\ WASHINGTON ST to JACKSON ST (700 - 799)",
			"address": "735 MONTGOMERY ST",
			"food_items": "Hot coffee: iced coffee: hot chocolate: tea: pastries",
			"latitude": 37.79621549659414,
			"longitude": -122.40375455824538,
			"schedules": ""
		},
		{
			"food_truck_name": "Wu Wei LLC dba MoBowl",
			"facility_type": "Truck",
			"location_description": "FRONT ST: VALLEJO ST to GREEN ST (900 - 999)",
			"address": "900 FRONT ST",
			"food_items": "Various types of meat: veggie: and seafood bowls.",
			"latitude": 37.80022091559222,
			"longitude": -122.39956901857713,
			"schedules": ""
		}
	]`
	var records []data.FoodTruckLocation
	_ = json.Unmarshal([]byte(text), &records)
	return records
}

type MockRepository struct{}

func (mr *MockRepository) BuildFoodTruckDataMap() error {
	return nil
}

func (mr *MockRepository) GetFoodTruckNearbyLocation(lat float64, lon float64, radiusMiles float64) ([]data.FoodTruckNearby, error) {
	return createMockFoodTrucksNearby(), nil
}

func (mr *MockRepository) GetFoodTruckLocation(latitude float64, longitude float64) (*data.FoodTruckLocation, error) {
	data := createMockFoodTrucksLocation()
	return &data[0], nil
}

func (mr *MockRepository) GetFoodTrucks() ([]data.FoodTruckLocation, error) {
	return createMockFoodTrucksLocation(), nil
}

type MockService struct {
	Repository repository.FoodTruckRepository
}

func (ms *MockService) GetFoodTruckNearbyLocation(lat float64, long float64, radiusMiles float64) ([]data.FoodTruckNearby, error) {
	return ms.Repository.GetFoodTruckNearbyLocation(lat, long, radiusMiles)
}

func (ms *MockService) GetFoodTruckLocation(latitude float64, longitude float64) (*data.FoodTruckLocation, error) {
	data, _ := ms.Repository.GetFoodTruckLocation(latitude, longitude)
	return data, nil
}

func (ms *MockService) GetFoodTrucks() ([]data.FoodTruckLocation, error) {
	return ms.Repository.GetFoodTrucks()
}
