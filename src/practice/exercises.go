package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//https://api.darksky.net/forecast/cb957c717f54f7a29bfb14de577110cc/-0.11982,51.51120

	// type Message struct {
	// 	Lat  float32
	// 	Long float32
	// }

	// var m Message

	// response, err := http.Get("https://api.darksky.net/forecast/cb957c717f54f7a29bfb14de577110cc/-0.11982,51.51120")
	// if err != nil {
	// 	fmt.Printf("Ups! The HTTP request to failed with error %s\n:", err)
	// } else {
	// 	data, _ := ioutil.ReadAll(response.Body)
	// 	var f interface{}
	// 	err := json.Unmarshal([]byte(data), &f)
	// 	if err != nil {
	// 		fmt.Println("error:", err)
	// 	}

	// 	fmt.Printf("%+v", f)
	// }

	//b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	response, err := http.Get("https://api.darksky.net/forecast/cb957c717f54f7a29bfb14de577110cc/-0.11982,51.51120")
	if err != nil {
		fmt.Printf("Ups! The HTTP request to failed with error %s\n:", err)
	} else {

		data, _ := ioutil.ReadAll(response.Body)
		var f interface{}
		err := json.Unmarshal(data, &f)
		if err != nil {
			fmt.Println("error", err)
		}

		//declare a type assertion to access f's underlying map
		m := f.(map[string]interface{})
		//iterate through the map with a range statement
		// for k, v := range m {
		// 	switch vv := v.(type) {
		// 	case string:
		// 		fmt.Println(k, " is string ", vv)
		// 	case float64:
		// 		fmt.Println(k, " is float64 ", vv)
		// 	case []interface{}:
		// 		fmt.Println(k, " is an array ", vv)
		// 		for i, u := range vv {
		// 			fmt.Println(i, u)
		// 		}
		// 	case map[string]interface{}:
		// 		fmt.Println(k, " is of type map ")
		// 		// if string(k) = "currently" {
		// 		// 	for key, value := range vv {
		// 		// 		fmt.Println("key: ", key, " value ", value)
		// 		// 	}
		// 		// }

		// 	default:
		// 		fmt.Printf(k, " is of type I can't handle %T")
		// 	}

		// }
		fmt.Println("lat: ", m["latitude"])
		currentW := m["currently"].(map[string]interface{})

		fmt.Println(">>>", currentW["temperature"])

		for k, v := range currentW {
			fmt.Println("k", k, " -- v -- ", v)
		}

	}
	//fmt.Println(">> ", f)

}
