package main

import (
	"testing"
	"fmt"
)

func TestGetWeather(t *testing.T){
	got := GetWeather()

	fmt.Println(got)

	
}