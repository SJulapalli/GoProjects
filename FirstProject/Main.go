package main

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

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

func submit(docName, firstName, lastName, email, message string) {
	opt := option.WithCredentialsFile("/Users/Suhas/GoProjects/dreamx-website-firebase-adminsdk-oc0et-5674edd2a0.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	client, err := app.Firestore(context.Background())

	if err != nil {
		log.Fatalln(fmt.Errorf("error initializing app: %v", err))
	}
	defer client.Close()

	result, err := client.Collection("Applicants").Doc(docName).Set(context.Background(), map[string]string{
		"first_name": firstName, "last_name": lastName, "email": email, "message": message,
	}, firestore.MergeAll)

	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)
}

func main() {
	submit("Suhas Julapalli", "Suhas", "Julapalli", "suhasj@terpmail.umd.edu", "This project has me very, very excited")
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
