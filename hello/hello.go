package main

import "fmt"

// func Hello() string {
// 	return "Hello World"
// }

func Hello(person_name string) string {
	return "Hello " + person_name
}

func main() {
	fmt.Println(Hello("Sriman"))
}
