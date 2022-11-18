package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "roma"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "NewDelhi"

	for country := range countryCapitalMap {
		fmt.Println(country, "capital ", countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["American"]
	if ok {
		fmt.Println("American capital ", capital)
	} else {
		fmt.Println("American capital not exists")
	}
}
