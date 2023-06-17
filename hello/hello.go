package main

import (
	"fmt"
	"greetings"
)

func main() {
	values := []string{"hyf", "hy"}
	message, err := greetings.Hellos(values)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(message)

}

func init() {
	fmt.Print("init")
}
