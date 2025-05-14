package main

import (
	"fmt"
)

func main() {
	greeting := greet("ja")
	fmt.Println(greeting)
}

type language string

// greet returns a greeting to the world
func greet(l language) string {
	switch l {
	case "fr":
		return "Bonjour le monde"
	case "es":
		return "Hola mundo"
	case "de":
		return "Hallo Welt"
	case "it":
		return "Ciao mondo"
	case "pt":
		return "Olá mundo"
	case "zh":
		return "你好，世界"
	case "ja":
		return "こんにちは世界"
	case "ru":
		return "Привет мир"
	case "ar":
		return "مرحبا بالعالم"
	default:
		return "Hello world"
	}
}
