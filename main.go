package main

import (
	"database/sql"
	"fmt"

	"github.com/jhonatanMota/go-starting/internal/infra/database"
	"github.com/jhonatanMota/go-starting/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	name := "Jhonatan Mota"

	println(Hello(name))

	car := Car{
		Model: "BMW",
		Color: "Blue",
	}

	car.Start()

	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)
	userCase := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "1",
		Price: 10.0,
		Tax:   2.0,
	}

	userCase.Execute(input)
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
