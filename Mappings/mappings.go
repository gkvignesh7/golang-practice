package main

import "fmt"

func main() {

	var Captialmapping map[string]string
	Captialmapping = make(map[string]string)
	Captialmapping["Tamil Nadu"] = "Chennai"
	Captialmapping["Goa"] = "Panaji"
	Captialmapping["Karnataka"] = "Bangalore"
	for country := range Captialmapping {
		fmt.Println( country,  Captialmapping[country])
		
	}
}