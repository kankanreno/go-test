package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@172.20.10.12:31235/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	deliveryChan, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	//forever := make(chan struct{})
	go func() {
		for delivery := range deliveryChan {
			fmt.Printf("Received a message: %s\n", delivery.Body)
		}
	}()
	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	//<- forever
	select{}
}
