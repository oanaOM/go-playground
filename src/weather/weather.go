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


func myRequestHTTP(url string) map[string]interface{}{
	var responseData map[string]interface{}
	response, err := http.Get(url)
	
	fmt.Println("Starting my app ... ")
	
	if err != nil {
		fmt.Printf("Ups! The HTTP request to " + url + " failed with error %s\n:", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		//initialise an empty interface to umarsh the JSON data
		json.Unmarshal([]byte(data), &responseData)
		fmt.Println("Request succesfully to " , url , " ... ")
    }
	//fmt.Println("My HTTP response data: ", responseData)

	return responseData
}


func main() {
	
	var getSingleIP, longitude, latitude string
	myGeoLocationKey := "d2dfba9048bf4c7594fc4c17f0b0956c"
	myDarkSkyKey := "cb957c717f54f7a29bfb14de577110cc"
	

	//start request to get IP
	getIPs := myRequestHTTP("https://httpbin.org/ip")
	//TODO: change to get everything before first , found
	runes :=[]rune(getIPs["origin"].(string))
	getSingleIP = string(runes[0:13])
	//fmt.Println("IP: ", getSingleIP)


	//start request to get lat and long
	getGeoLocation := myRequestHTTP("https://api.ipgeolocation.io/ipgeo?apiKey=" + myGeoLocationKey + "&ip=" + getSingleIP)
	longitude = fmt.Sprint(getGeoLocation["longitude"])
	latitude = fmt.Sprint(getGeoLocation["latitude"])
	//fmt.Println("Your latitude: ", latitude, " and longitude ", longitude, " city", fmt.Sprint(getGeoLocation["city"]))
	

	//start request to get lat and long
	getForecast := myRequestHTTP("https://api.darksky.net/forecast/" + myDarkSkyKey + "/" + longitude + "," + latitude)
	weatherToday := getForecast["hourly"].(map[string]interface{})
	// weatherTodayData := weatherToday["data"].(map[string]interface{})
	//fmt.Println("weatherToday: ", weatherToday)

	//TODO: convert temperature to Celsius
	fmt.Println("Having a nice time in", fmt.Sprint(getGeoLocation["city"], " ?"), 
				" \n It looks like the weather's ", weatherToday["summary"])	
				
	// fmt.Println("Having a nice time in", fmt.Sprint(getGeoLocation["city"], " ?"), 
	// 			" \n It looks like the weather's ", weatherToday["summary"],
	// 			"\n with a temperature of ", weatherToday["data"], "Fahrenheit" )

}
