package main

import "fmt"


// Also, we can group constants in a block instead of declaring them each on their own line. 
// It's a good idea to use a line between sets of related constants for readability.
const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)


// The function name starts with a uppercase letter. 
// In Go, public functions start with a capital letter and private ones start with a lowercase.
func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// note the named return value prefix
//   - It will be assigned the "zero" value. This depends on the type, for example ints are 0 and for strings it is "".
//   - You can return whatever it's set to by just calling return rather than return prefix.
//   - This will display in the Go Doc for your function so it can make the intent of your code clearer.
func greetingPrefix(language string) (prefix string) { 
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		// default in the switch case will be branched to if none of the other case statements match.
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
