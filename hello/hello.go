package main

import (
	"fmt"

	"github.com/mysticparrot/golang-learning/greetings"
)

func main() {
	message := greetings.Hello("Gladdy's")
	fmt.Printf(message)
}
