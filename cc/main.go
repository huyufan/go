package main

import (
	"context"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	fmt.Println(ctx)
	cancel()

	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	cancel1()
	fmt.Println(ctx1)
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	_, err = conn.Do("lpush", "book_list", "abc", "ceg")
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.String(conn.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(r)

}
