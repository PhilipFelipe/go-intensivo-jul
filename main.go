package main

import (
	"database/sql"
	"fmt"

	"github.com/PhilipFelipe/go-intensivo-jul/internal/infra/database"
	"github.com/PhilipFelipe/go-intensivo-jul/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
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
	input := usecase.OrderInput{
		ID:    "1234",
		Price: 10.0,
		Tax:   1.0,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
