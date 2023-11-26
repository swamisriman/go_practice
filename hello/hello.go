package main

import "fmt"

const helloPrefix = "Hello "

func Hello(personName string) string {
	if personName == "" {
		return helloPrefix + "World"
	}
	return helloPrefix + personName
}

func main() {
	fmt.Println(Hello("Sriman"))
}
