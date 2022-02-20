package main

import (
    "fmt"

    "github.com/Minjia-theodore-Chen/first_go_app/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}