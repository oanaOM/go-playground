package main

/*
* uses accuWeather API



import all necessary libraries below
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Starting my application ... ")
	//get the user IP
	var getIPs, getSingleIP string
	//start HTTP request to get IP
	response, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		fmt.Printf("Ups! Couldn't get IP. The HTTP request failed with error %s\n:", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		//initialise an empty interface to umarsh the JSON data
		var jsonData map[string]interface{}
		json.Unmarshal([]byte(data), &jsonData)

		getIPs = jsonData["origin"].(string)
		runes := []rune(getIPs)
		getSingleIP := string(runes[0:14])

		fmt.Println("IP: ", getSingleIP)

	}

	//unique key to access accuWeather
	myAccessKey := "jGLWMiB6XGtJ0hJZMCSzBGsYu8T0GJx9"
	//initialise a keylocation var to construct the request for 1 day forecast
	var keyLocation string

	//get location key
	// http://dataservice.accuweather.com/locations/v1/cities/ipaddress?apikey=jGLWMiB6XGtJ0hJZMCSzBGsYu8T0GJx9&q=86.147.198.233
	response, err = http.Get("http://dataservice.accuweather.com/locations/v1/cities/ipaddress?apikey=" + myAccessKey + "&q=" + getSingleIP)

	if err != nil {
		fmt.Printf("Ups! Couldn't get the location key. The HTTP request failed with error %s\n: ", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println("\n", string(data))

		//unmarch the JSON using an empty interface
		var jsonData map[string]interface{}
		//unmarch/decode
		json.Unmarshal([]byte(data), &jsonData)

		keyLocation := jsonData["Key"]

		fmt.Println(">>>> LocalizedName: ", jsonData["LocalizedName"], " and the key: ", keyLocation)
	}

	/* 1 day forecast request*/
	// http://dataservice.accuweather.com/forecasts/v1/daily/1day/328328?apikey=jGLWMiB6XGtJ0hJZMCSzBGsYu8T0GJx9

	response, err = http.Get("http://dataservice.accuweather.com/forecasts/v1/daily/1day/" + keyLocation + "?apikey=" + myAccessKey)

	if err != nil {
		fmt.Printf("Ups! Couldn't get the 1 day forecast. The HTTP request failed with error %s\n: ", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println("\n", string(data))

		//unmarch the JSON using an empty interface
		var jsonData map[string]interface{}
		//unmarch/decode
		json.Unmarshal([]byte(data), &jsonData)

		//headline := jsonData["Headline"].(map[string]interface{})

		fmt.Println(">>>> headline: ", jsonData["Headline"])
	}

}
