package main

import "fmt"

func main() {
	countryCapitalMap := map[string]string{"France": "Paris",
		"Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("raw map")

	for country := range countryCapitalMap {
		fmt.Println(country, "capital", countryCapitalMap[country])
	}

	delete(countryCapitalMap, "France")
	fmt.Println("France has been deleted")

	for country := range countryCapitalMap {
		fmt.Println(country, "capital", countryCapitalMap[country])
	}
}
