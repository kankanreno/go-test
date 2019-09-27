package main

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/streadway/amqp"
	"time"
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

	for i := 0; ; i++ {
		body := cast.ToString(i) + " Hello World!"
		if err := ch.Publish("", queue.Name, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		}); err != nil {
			panic(err)
		}
		fmt.Printf(" [x] Sent %s\n", body)

		time.Sleep(1000 * time.Millisecond)
	}
}
