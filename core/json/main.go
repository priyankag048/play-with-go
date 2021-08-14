package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Marshalling JSON: convert Go object to json string
	type Book struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	book := Book{
		Title:  "Learning Concurreny in Python",
		Author: "Elliot Forbes",
	}
	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))

	// Unmarshalling
	type SensorReading struct {
		Name     string `json:"name"`
		Capacity int    `json:"capacity"`
		Time     string `json:"time"`
	}
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`
	var reading SensorReading
	error := json.Unmarshal([]byte(jsonString), &reading)
	fmt.Printf("%+v\n", reading)
}
