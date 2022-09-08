package main

import (
	"fmt"
	"runtime"

	"google.golang.org/api/option"
)

type Car struct {
	Brand      string `required max: "100"`
	Color      string
	Seats      int
	Passengers []string
}

func (c Car) print() {
	fmt.Println(c)
}

type Honda interface {
	drive() string
	park() bool
}

type Prius struct {
}

func (p Prius) drive() string {
	return "Vroom"
}

func (p Prius) park() bool {
	return true
}

func main() {
	sa := option.WithCredentialsFile("./ServiceAccountKey.json")
	print(sa)
	grades := []int{1, 2, 3, 4, 5}
	b := make([]int, 3, 100)
	valMap := map[string]int{
		"Suhas":  19,
		"Subas":  18,
		"Sinhas": 17,
	}

	pop, ok := valMap["Suhas"]

	honda := Car{
		Brand:      "honda",
		Color:      "blue",
		Seats:      5,
		Passengers: []string{"John, Alex, Charles"},
	}
	defer fmt.Println("Defered", valMap, pop, ok, len(valMap))
	fmt.Println(cap(grades) + len(b))
	fmt.Println(honda)
	fmt.Println(struct {
		Car
		name string
	}{Car: Car{Brand: "Toyota", Color: "Green", Seats: 2, Passengers: []string{"Ripley"}}, name: "Ripley"})

	if pop, ok := valMap["Suhas"]; ok {
		fmt.Println(pop)
	}

	switch pop := valMap["Suhas"]; pop {
	case 19:
		fmt.Println("Suhas")
	default:
		fmt.Println("Unknown")
	}

	num := Car{
		Brand:      "honda",
		Color:      "Green",
		Seats:      5,
		Passengers: []string{"Andrew", "Josh"},
	}

	num2 := &num
	fmt.Println(num, num2.Brand)

Loop:
	for k, v := range valMap {
		fmt.Println(k, v)
		if k == "Sinhas" {
			break Loop
		}
	}

	honda.print()

	var x Honda = Prius{}

	fmt.Println(x.drive())
	fmt.Println(runtime.GOMAXPROCS(-1))
}
