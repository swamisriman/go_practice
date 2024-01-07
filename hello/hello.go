package main

import "fmt"

const (
	english = "English"
	spanish = "Spanish"
	telugu  = "Telugu"

	englishGreeting = "Hello"
	spanishGreeting = "Hola"
	teluguGreeting  = "నమస్తే"
)

func Hello(lang, personName string) string {
	if personName == "" {
		personName = "World"
	}

	return GetGreetingInLang(lang) + " " + personName
}

func GetGreetingInLang(lang string) string {
	switch lang {
	case telugu:
		return teluguGreeting
	case english:
		return englishGreeting
	case spanish:
		return spanishGreeting
	default:
		return englishGreeting
	}
}

func main() {
	fmt.Println(Hello("English", "Sriman"))
}
