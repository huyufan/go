package main

import (
	"fmt"
)

func main() {
	sd := new([]int)
	fmt.Println(sd)
	fmt.Println(&sd)
	news := []int{1, 2, 3}
	sd = &news
	fmt.Println(&sd)
	fmt.Println(&news)
	fmt.Println(*sd)
	fmt.Println(news)

	list := make([]int, 0)
	list = append(list, 1)
	fmt.Println(list)
	// rabbitmq := RabbitMQ.NewRabbitMQPubSub("" +
	// 	"newProduct")
	//rabbitmq.RecieveSub()
}
