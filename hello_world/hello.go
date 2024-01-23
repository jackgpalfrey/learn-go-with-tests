package main

import "fmt"

const (
	LangEnglish = "English"
	LangSpanish = "Spanish"
	LangFrench  = "French"
)

const (
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishPrefix
	switch lang {
	case LangSpanish:
		prefix = spanishPrefix

	case LangFrench:
		prefix = frenchPrefix
	}

	return prefix + name
}

func main() {
	fmt.Printf(Hello("Jack", LangEnglish))
}
