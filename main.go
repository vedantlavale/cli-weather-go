package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				Time_Epoch int64   `json:"time_epoch"`
				Time       string  `json:"time"`
				TempC      float64 `json:"temp_c"`
				Condition  struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

// first define the main function, and add the error checking logic
// then add the HTTP GET request to fetch weather data from the API
// finally, read the response body and print it as a string
func main() {
	loc := "Pune" // You can change this to any location you want to query
	if(len(os.Args)>=2){
		loc=os.Args[1]
	}

	fmt.Print("Please enter a location: ")
	fmt.Scanln(&loc)
	err := godotenv.Load()   // Load the .env file to get the API key
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")  // Get the API key from the environment variable
	if apiKey == "" {
		log.Fatal("API key not set")
	}
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key="+apiKey+"&q="+loc+"&days=1&aqi=no&alerts=no")

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Please check the location and try again. " + res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather

	err = json.Unmarshal(body, &weather) // Assign the result of json.Unmarshal to err
	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf("%s,%s: %.0fC,%s \n",
		location.Name, 
		location.Country, 
		current.TempC, 
		current.Condition.Text,
	)

	for _,hour := range hours{
		
		date:=time.Unix(hour.Time_Epoch,0)

		if date.Before(time.Now()) {
			continue // Skip past hours
		}

			message:=fmt.Sprintf(
				"%s - %.0fC, %.0f%%, %s\n",
				date.Format("15:04"),
				hour.TempC,
				hour.ChanceOfRain,
				hour.Condition.Text,
			)

			if(hour.ChanceOfRain < 40){
			fmt.Print(message)
			} else{
				color.Red(message)
			}
		}
		
	}
