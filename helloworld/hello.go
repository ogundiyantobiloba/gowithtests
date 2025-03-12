package helloworld

import "fmt"

func main() {
	fmt.Println(Hello("chris", ""))
}

const (
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "

	spanish = "Spanish"
	french  = "French"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishHelloPrefix

	case french:
		prefix = frenchHelloPrefix

	default:
		prefix = englishHelloPrefix
	}
	return
}
