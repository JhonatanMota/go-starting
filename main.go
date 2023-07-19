package main

import (
	"fmt"
)

func main() {
	name := "Jhonatan Mota"

	println(Hello(name))

	car := Car{
		Model: "BMW",
		Color: "Blue",
	}

	car.Start()
}

func Hello(name string) string {
	message := fmt.Sprintf("Hello %v. Welcome to my learning go journey!", name)
	return message
}

type Car struct {
	Model string
	Color string
}

// method
func (c Car) Start() {
	println(c.Model + " is starting")
}

// function
func sum(x, y int) int {
	return x + y
}
