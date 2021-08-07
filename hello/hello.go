package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, lang string) string {
	prefix := englishHelloPrefix
	if lang == "Spanish" {
		prefix = spanishHelloPrefix
	}
	if len(name) == 0 {
		name = "world"
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
