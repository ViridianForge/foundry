package hello

import "fmt"

const engHello = "Hello, "
const spanishHello = "Hola, "
const frenchHello = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return helloPrefix(language) + name + "!"
}

func helloPrefix(language string) string {
	switch language {
	case "Spanish":
		return spanishHello
	case "French":
		return frenchHello
	default:
		return engHello
	}
}

func main() {
	fmt.Println(Hello("World", "English"))
}
