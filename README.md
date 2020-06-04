# Weather or Not!
A simple weather app that shows the weather at your current location.

## Overview
This app will retrieve the user's IP and feed it to a [P Geolocation API.](https://ipgeolocation.io/) request to retrieve the latitude and longitude coordinates used by [Dark Sky API ](https://darksky.net/dev)  to show the local weather.

## Features

 - [x] Retrieve user's location
 - [x] Retrieve local time temperature 
 - [x] Retrieve a short weather summary
 - [x] Show everything in UI (probably html page)

# Getting started

## Installation  
For this app to work, you need to have [Go](https://golang.org) installed on your machine. Please follow the instructions on their site on how to do this.


## How to run weather web app locally
Once you have Go up and running, open a terminal, clone this repo in a new folder, navigate to  *src/webApp* folder and type 

```
cd src

go run main.go
```

Open your favorite browser and navigate to your http://127.0.0.1:8080/

## Docker
Run the app in a Docker container as such
```
docker image build -t myweatherapp .

docker container run -p 9991:8080 myweatherapp
```


------

### References

- Official Golang  - https://golang.org/ 
- Building a Golang Docker image - https://bitfieldconsulting.com/golang/docker-image

 