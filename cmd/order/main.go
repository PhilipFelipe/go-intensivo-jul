package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/PhilipFelipe/go-intensivo-jul/internal/infra/database"
	"github.com/PhilipFelipe/go-intensivo-jul/internal/usecase"
	"github.com/PhilipFelipe/go-intensivo-jul/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Car struct {
	Model string
	Color string
}

// metodo
func (c Car) Start() {
	println(c.Model + " starting...")
}

func (c *Car) ChangeColor(color string) {
	// Quando NÃO tem asterisco duplica o valor de c.color na memória - cópia do color original
	c.Color = color
	println("New color:", c.Color)
}

// funcao
func soma(x, y int) int {
	return x + y
}

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	msgRabbitmqChan := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChan) // T2
	rabbitmqWorker(msgRabbitmqChan, uc)      // T1
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq")
	for msg := range msgChan {
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
		fmt.Println("Mensagem processada e salva no banco:", output)
	}
}
