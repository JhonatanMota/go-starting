package main

import (
	"fmt"
)

func main() {
	name := "Jhonatan Mota"

	println(Hello(name))
}

func Hello(name string) string {
	message := fmt.Sprintf("Hello %v. Welcome to my learning go journey!", name)
	return message
}
