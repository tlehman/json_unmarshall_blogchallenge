package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name string
	Age  float64
}

type Place struct {
	City    string
	Country string
}

type Mixed struct {
	Name    string
	Age     float64
	City    string
	Country string
}

func solutionA(jsonString []byte) ([]Person, []Place) {
	persons := []Person{}
	places := []Place{}

	// data is a map from strings to (an array of maps from strings to interface)
	var data map[string][]map[string]interface{}
	err := json.Unmarshal(jsonString, &data)

	// jsonString may not be a proper serialized json string
	if err != nil {
		fmt.Println(err)
		return persons, places
	}

	for i := range data["things"] {
		item := data["things"][i]

		if item["name"] != nil {
			persons = addPerson(persons, item)
		} else {
			places = addPlace(places, item)
		}
	}

	return persons, places
}

func solutionB(jsonString []byte) ([]Person, []Place) {
	persons := []Person{}
	places := []Place{}

	var data map[string][]Mixed

	err := json.Unmarshal(jsonString, &data)
	if err != nil {
		fmt.Println(err)
		return persons, places
	}

	for i := range data["things"] {
		item := data["things"][i]
		if item.Name != "" {
			persons = append(persons, Person{item.Name, item.Age})
		} else {
			places = append(places, Place{item.City, item.Country})
		}
	}

	return persons, places
}

func addPerson(persons []Person, item map[string]interface{}) []Person {
	name, _ := item["name"].(string)
	age, ok := item["age"].(float64)
	if !ok {
		fmt.Println("age = ", age)
	}
	person := Person{name, age}
	persons = append(persons, person)
	return persons
}

func addPlace(places []Place, item map[string]interface{}) []Place {
	city, _ := item["city"].(string)
	country, _ := item["country"].(string)
	place := Place{city, country}
	places = append(places, place)
	return places
}

func solutionC(jsonStr []byte) ([]Person, []Place) {
	persons := []Person{}
	places := []Place{}

	var data map[string][]json.RawMessage

	err := json.Unmarshal(jsonStr, &data)

	if err != nil {
		fmt.Println(err)
		return persons, places
	}
	for _, thing := range data["things"] {
		persons = addPersonC(thing, persons)
		places = addPlaceC(thing, places)
	}
	return persons, places
}

func addPersonC(thing json.RawMessage, persons []Person) []Person {
	person := Person{}
	err := json.Unmarshal(thing, &person)
	if err != nil {
		fmt.Println(err)
	} else {
		if person != *new(Person) { // new(Person) allocates and zeros a Person, and returns a pointer
			persons = append(persons, person)
		}
	}
	return persons
}

func addPlaceC(thing json.RawMessage, places []Place) []Place {
	place := Place{}
	err := json.Unmarshal(thing, &place)
	if err != nil {
		fmt.Println(err)
	} else {
		if place != *new(Place) { // new(Place) allocates and zeros a Place, and returns a pointer
			places = append(places, place)
		}
	}

	return places
}

func main() {
	data, err := ioutil.ReadFile("people_places.json")

	if err != nil {
		fmt.Println(err)
	}

	personsA, placesA := solutionA(data)
	personsB, placesB := solutionB(data)
	personsC, placesC := solutionC(data)

	fmt.Println(personsA, placesA)
	fmt.Println("\n")
	fmt.Println(personsB, placesB)
	fmt.Println("\n")
	fmt.Println(personsC, placesC)
}
