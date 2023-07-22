package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/jhonatanMota/go-starting/internal/infra/database"
	"github.com/jhonatanMota/go-starting/internal/usecase"
	"github.com/jhonatanMota/go-starting/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"
	ampq "github.com/rabbitmq/amqp091-go"
)

func main() {

	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	defer db.Close() // after everything is done close the database connection

	orderRepository := database.NewOrderRepository(db)
	userCase := usecase.NewCalculateFinalPrice(orderRepository)

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close() // after everything is done close the rabbitmq connection is closed

	msgRabbitmqChannel := make(chan ampq.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel)
	rabbitmqWorker(msgRabbitmqChannel, userCase)

}

func rabbitmqWorker(msgChan chan ampq.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("rabbitmqWorker started")
	for msg := range msgChan {
		// deserializing json to struct
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}

		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println("Message received and saved on database", output)
	}
}
