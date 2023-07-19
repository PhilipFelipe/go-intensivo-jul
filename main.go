package main

import (
	"fmt"
	"time"
)

func process() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

// T1
func main() {

	canal := make(chan int)

	// T2
	go func() {
		for i := 0; i < 10; i++ {
			canal <- i // envia dados para o canal
			fmt.Println("Jogou no canal:", i)
		}
	}()
	// T3
	go func() {
		for i := 0; i < 10; i++ {
			canal <- i // envia dados para o canal
			fmt.Println("Jogou no canal:", i)
		}
	}()

	// T4
	go worker(canal, 1)
	go worker(canal, 2)
	go worker(canal, 3)
	worker(canal, 4)
}

func worker(canal chan int, workerID int) {
	for {
		fmt.Println("Recebeu do canal:", <-canal, "no worker:", workerID)
		time.Sleep(time.Second)
	}
}
