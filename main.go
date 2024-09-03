package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	fmt.Println("Successfully connected to RabbitMQ!")
}
