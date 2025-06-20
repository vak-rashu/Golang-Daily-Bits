package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var api_key string = os.Getenv("API_KEY")

type WeatherJson struct {
	Location struct {
		Country  string `json:"country"`
		TimeZone string `json:"timezone_id"`
	} `json:"location"`
	CurrentTemp struct {
		Obs_Time    string   `json:"observation_time"`
		Temp        int      `json:"temperature"`
		WeatherDesc []string `json:"weather_descriptions"`
		AirQuality  struct {
			CO  string `json:"co"`
			O3  string `json:"o3"`
			So2 string `json:"so2"`
		} `json:"air_quality"`
		WindSpeed int `json:"wind_speed"`
	} `json:"current"`
	Success string `json:"success"`
	IdError struct {
		Code int    `json:"code"`
		Type string `json:"type"`
		Info string `json:"info"`
	} `json:"error"`
}

var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Weather API CLI",
	Long:  `Weather API CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		country, _ := cmd.Flags().GetString("country")
		weatherApi(country)
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)
	weatherCmd.Flags().StringP("country", "c", "", "Country name")
}

func weatherApi(country string) {
	req := fmt.Sprintf("http://api.weatherstack.com/current?access_key=%v&query="+country, api_key)
	res, err := http.Get(req)
	if err != nil {
		fmt.Println("Error occurred, hence could not proceed further:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error occured in reading body:", err)
		return
	}

	var weather WeatherJson
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println(err)
		return
	}

	location := fmt.Sprintf("Country: %s\nTime Zone: %s\n", weather.Location.Country, weather.Location.TimeZone)
	currentTemp := fmt.Sprintf("Observation Time: %s\nTemperature: %v\nWeather Description: %v\n", weather.CurrentTemp.Obs_Time, weather.CurrentTemp.Temp, weather.CurrentTemp.WeatherDesc)
	airQuality := fmt.Sprintf("Air Quality Check:\nco- %s\no3 - %s\nSo2 - %s", weather.CurrentTemp.AirQuality.CO, weather.CurrentTemp.AirQuality.O3, weather.CurrentTemp.AirQuality.So2)
	fmt.Println(location)
	fmt.Println(strings.Repeat("*", 50))
	fmt.Println(currentTemp)
	fmt.Println(strings.Repeat("*", 50))
	fmt.Println(airQuality)
	fmt.Println(strings.Repeat("*", 50))
}
