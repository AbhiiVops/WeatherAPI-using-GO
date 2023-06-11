package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Weather struct {
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
	// Read API key from environment variable
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY environment variable not set")
	}

	// Read city parameter from the request URL
	params := mux.Vars(r)
	city := params["city"]

	// Make a request to the weather API
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("q", city).
		SetQueryParam("appid", apiKey).
		Get("https://api.openweathermap.org/data/2.5/weather")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the response JSON
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
	}
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		log.Fatal(err)
	}

	// Create a Weather struct from the parsed data
	weather := Weather{
		Temperature: data.Main.Temp,
		Description: data.Weather[0].Description,
	}

	// Send the weather data as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load("weather_api.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new router
	router := mux.NewRouter()

	// Define the API routes
	router.HandleFunc("/weather/{city}", GetWeather).Methods("GET")

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
