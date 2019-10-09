package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number int
	actorName string
	companions []string
}

type Animal struct {
	Name string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly bool
}

// Tags
type Car struct {
	Name string `required max:"100"`
	Model string
}

func main() {
	// Map
	statePopulations := map[string]int{
		"Lagos": 39250017,
		"Abia": 27862596,
		"Oyo": 20612439,
		"Kano": 19745289,
		"Anambra": 12802503,
		"Zamfara": 12801539,
		"Benue": 11614373,
	}
	bayelsa, ok := statePopulations["Bayelsa"] 
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Lagos"])
	fmt.Println(bayelsa, ok)

	aDoctor := Doctor{
		number: 3,
		actorName: "John Doe",
		companions: []string {
			"Jane Doe",
			"Johnny English",
			"Baba Suwe",
		},
	}
	fmt.Println(aDoctor)

	bDoctor := struct{name string}{name: "Red Fox"} // anonymous struct
	fmt.Println(bDoctor)

	// Composition
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b.Name)

	// Tags
	t := reflect.TypeOf(Car{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}