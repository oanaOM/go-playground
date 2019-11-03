package main

/*
* Dark Sky API
* ipgeolocation
*
*
*
* NoteToSelf: use google mails
 */

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func myRequestHTTP(url string) map[string]interface{} {
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
	//fmt.Println("My HTTP response data: ", responseData)

	return responseData
}

func main() {

	var longitude, latitude string
	myGeoLocationKey := "d2dfba9048bf4c7594fc4c17f0b0956c"
	myDarkSkyKey := "cb957c717f54f7a29bfb14de577110cc"

	//start request to get lat and long
	getGeoLocation := myRequestHTTP("https://api.ipgeolocation.io/ipgeo?apiKey=" + myGeoLocationKey)
	longitude = fmt.Sprint(getGeoLocation["longitude"])
	latitude = fmt.Sprint(getGeoLocation["latitude"])

	//start request to get lat and long
	getForecast := myRequestHTTP("https://api.darksky.net/forecast/" + myDarkSkyKey + "/" + latitude + "," + longitude)
	weatherToday := getForecast["currently"].(map[string]interface{})
	tempF := weatherToday["temperature"].(float64)
	tempC := (tempF - float64(32)) * float64(5) / float64(9)

	//TODO: convert temperature to Celsius
	fmt.Println("Having a nice time in --- ", fmt.Sprint(getGeoLocation["city"], " ?"),
		"\nToday's weather is    --- ", weatherToday["summary"])
	fmt.Printf("with a temperature of ---  %.0f", tempC)
	fmt.Println(" C")

}
