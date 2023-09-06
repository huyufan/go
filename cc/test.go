package main

import "fmt"

// Greeting function types
type Greeting func(name string) string

func (g Greeting) say(n string) {
	fmt.Println(g(n))
}

func english(name string) string {
	return "Hello, " + name
}

func mainss() {
	greet := Greeting(english)
	//fmt.Println(greet)
	greet.say("World")
}
