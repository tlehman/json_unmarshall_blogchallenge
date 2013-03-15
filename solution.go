package main

import (
  "fmt"
  "io/ioutil"
  "encoding/json"
)

type Person struct {
  Name string
  Age  int
}

type Place struct {
  City    string
  Country string
}

func solutionA(jsonString []byte) ([]Person, []Place) {
  persons := []Person{}
  places := []Place{}

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

func addPerson(persons []Person, item map[string]interface{}) ([]Person) {
  name, _ := item["name"].(string)
  age, _ := item["age"].(int)
  person := Person{name,age}
  persons = append(persons, person)
  return persons
}

func addPlace (places []Place, item map[string]interface{}) ([]Place){
  city, _ := item["city"].(string)
  country, _ := item["country"].(string)
  place := Place{city,country}
  places = append(places, place)
  return places
}

func main() {
  data, err := ioutil.ReadFile("people_places.json")

  if err != nil {
    fmt.Println(err)
  }

  persons, places := solutionA(data)

  fmt.Println(persons)
  fmt.Println("\n")
  fmt.Println(places)
}


