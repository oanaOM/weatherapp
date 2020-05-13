package main
/*
* Weather or Not! is a weather app that shows the weather at your current location
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"text/template"
	"path"
	"log"
)

/*Weather custom type that stores our weather variables */
type Weather struct {
	City        string
	Summary     string
	Temperatura float64
	Icon 		string
}

func main() {

	http.HandleFunc("/", ShowWeather)
	if err:=http.ListenAndServe(":8080", nil);err!=nil{
		log.Fatalf("Could not listen on mentioned port: %v", err)
	}

}
// MyRequestHTTP retrieve the data
func MyRequestHTTP(url string) map[string]interface{} {
	var responseData map[string]interface{}
	response, err := http.Get(url)

	fmt.Println("Starting my app ... ")

	if err != nil {
		fmt.Printf("Ups! The HTTP request to "+url+" failed with error %s\n:", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		//initialise an empty interface for my data
		jsonErr := json.Unmarshal(data, &responseData)
		if jsonErr != nil {
			fmt.Println("error:", jsonErr)
		}
		fmt.Println("Request succesfully to ", url, " ... ")

	}
	return responseData
}
/*GetWeather retrieves the weather data */
func GetWeather() Weather {
	var longitude, latitude string
	myGeoLocationKey := "d2dfba9048bf4c7594fc4c17f0b0956c"
	myDarkSkyKey := "cb957c717f54f7a29bfb14de577110cc"

	//start request to get lat and long
	getGeoLocation := MyRequestHTTP("https://api.ipgeolocation.io/ipgeo?apiKey=" + myGeoLocationKey)
	longitude = fmt.Sprint(getGeoLocation["longitude"])
	latitude = fmt.Sprint(getGeoLocation["latitude"])

	//start request to get lat and long
	getForecast := MyRequestHTTP("https://api.darksky.net/forecast/" + myDarkSkyKey + "/" + latitude + "," + longitude)
	weatherNow := getForecast["currently"].(map[string]interface{})
	weatherToday := getForecast["daily"].(map[string]interface{})
	tempF := weatherNow["temperature"].(float64)
	tempC := (tempF - float64(32)) * float64(5) / float64(9)
	forecastData := Weather{fmt.Sprint(getGeoLocation["city"]), fmt.Sprint(weatherToday["summary"]), math.Round(tempC), fmt.Sprint(weatherNow["icon"])}

	return forecastData
}

/*ShowWeather is a handler func that binds the response with the HTML template */
func ShowWeather(w http.ResponseWriter, r *http.Request) {

	myWeather := GetWeather()
	
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, myWeather); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	w.Header().Set("Content-Type", "application/json")

}