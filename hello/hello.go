package main

import (
    "fmt"

    "github.com/minjia-theodore-mhen/first_go_app/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}