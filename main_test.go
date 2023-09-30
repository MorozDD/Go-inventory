package main

import (
	"bytes"
	"encoding/json"
	"inventory/routes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"inventory/api/handlers"
	"inventory/database"
	"inventory/models"

	"github.com/gin-gonic/gin"
)

var mainRouter *gin.Engine

func init() {
	// Initialize the database connection for testing
	_, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	// Create a new Gin router without starting it
	r := gin.New()
	routes.InitializeRoutes(r)

	// Replace the Gin router in the main function with the test router
	mainRouter = r
}

func TestCreatePart(t *testing.T) {
	part := models.Part{
		Name:        "Test Part",
		SKU:         "TEST123",
		Description: "Test Description",
		Price:       10.99,
		Location:    "Test Location",
		ShipmentPackaging: models.ShipmentPackaging{
			Weight:    1.5,
			Size:      "Small",
			Hazardous: false,
			Fragile:   true,
		},
		Attributes: []models.Attribute{
			{
				Name:  "Color",
				Value: "Red",
			},
			{
				Name:  "Size",
				Value: "Medium",
			},
		},
		Fitments: []models.Fitment{
			{
				Make:     "Toyota",
				CarModel: "Camry",
				Year:     2022,
			},
			{
				Make:     "Honda",
				CarModel: "Civic",
				Year:     2023,
			},
		},
		Images: []models.Image{
			{
				ImageURL: "https://example.com/image1.jpg",
			},
			{
				ImageURL: "https://example.com/image2.jpg",
			},
		},
		Metadata: []models.Metadata{
			{
				Key:   "Key1",
				Value: "Value1",
			},
			{
				Key:   "Key2",
				Value: "Value2",
			},
		},
	}

	partJSON, err := json.Marshal(part)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	req, _ := http.NewRequest("POST", "/part", bytes.NewBuffer(partJSON))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Perform the HTTP request
	mainRouter.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}
}

func TestGetPartByID(t *testing.T) {
	part := models.Part{
		Name:              "Test Part",
		SKU:               "TEST123",
		Description:       "Test Description",
		Price:             10.99,
		Location:          "Test Location",
		ShipmentPackaging: models.ShipmentPackaging{Weight: 1.0},
	}

	// Insert the test part into the database
	if err := handlers.CreatePart(&part); err != nil {
		t.Fatalf("Error creating test part: %v", err)
	}

	req, _ := http.NewRequest("GET", "/part/"+strconv.Itoa(int(part.ID)), nil)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Perform the HTTP request
	mainRouter.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestUpdatePart(t *testing.T) {
	part := models.Part{
		Name:              "Test Part",
		SKU:               "TEST123",
		Description:       "Test Description",
		Price:             10.99,
		Location:          "Test Location",
		ShipmentPackaging: models.ShipmentPackaging{Weight: 1.0},
	}

	// Insert the test part into the database
	if err := handlers.CreatePart(&part); err != nil {
		t.Fatalf("Error creating test part: %v", err)
	}

	// Modify the part
	part.Name = "Updated Part Name"

	partJSON, err := json.Marshal(part)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	req, _ := http.NewRequest("PUT", "/part/"+strconv.Itoa(int(part.ID)), bytes.NewBuffer(partJSON))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Perform the HTTP request
	mainRouter.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

}

func TestDeletePart(t *testing.T) {
	part := models.Part{
		Name:              "Test Part",
		SKU:               "TEST123",
		Description:       "Test Description",
		Price:             10.99,
		Location:          "Test Location",
		ShipmentPackaging: models.ShipmentPackaging{Weight: 1.0},
	}

	// Insert the test part into the database
	if err := handlers.CreatePart(&part); err != nil {
		t.Fatalf("Error creating test part: %v", err)
	}

	req, _ := http.NewRequest("DELETE", "/part/"+strconv.Itoa(int(part.ID)), nil)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Perform the HTTP request
	mainRouter.ServeHTTP(w, req)

	// Check the response status code and body
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetPartByVersion(t *testing.T) {
	part := models.Part{
		Name:              "Test Part321",
		SKU:               "TEST321",
		Description:       "Test Description321",
		Price:             20.99,
		Location:          "Test Location",
		ShipmentPackaging: models.ShipmentPackaging{Weight: 5.0},
		Version:           1, // Set the version
	}

	if err := handlers.CreatePart(&part); err != nil {
		t.Fatalf("Error creating part: %v", err)
	}

	retrievedPart, err := handlers.GetPartByVersion(part.ID, part.Version)
	if err != nil {
		t.Fatalf("Error getting part by version: %v", err)
	}

	// Check if the retrieved part matches the created part
	if retrievedPart.ID != part.ID {
		t.Errorf("Expected part ID %d, got %d", part.ID, retrievedPart.ID)
	}
	if retrievedPart.Version != part.Version {
		t.Errorf("Expected part version %d, got %d", part.Version, retrievedPart.Version)
	}
}
