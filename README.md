# Weather or Not!
A simple weather app that shows 1 day forecast at your current location.

## Overview
This app will retrieve the user's IP and feed it to a [P Geolocation API.](https://ipgeolocation.io/) request to retrieve the latitude and longitude coordinates used by [Dark Sky API ](https://darksky.net/dev)  to show the local weather.

## Features

 - [x] Retrieve user's location
 - [ ] Retrieve local time temperature 
 - [x] Retrieve a short weather summary
 - [ ] Retrieve pollution level and maybe a small warning
 - [ ] Show everything in UI (probably html page)
 - [ ] Change icon based on weather
 - [ ] Look into GITHub apps

# Getting started

## Installation and running 
Currently in order for this app to work, you need to have [Go](https://golang.org) installed on your machine. Please follow the instructions on their site on how to do this.
Once you have Go up and running, open a terminal,  clone this repo on your machine, navigate to  *bin/weather* folder and type 

    go run weather.go


 