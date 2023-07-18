package iteration

import "strings"

// func Repeat(character string, iterations int) string {
// 	var repeated string               // declare variable, := is a shortcut to declare and assign a value to a variable
// 	for i := 0; i < iterations; i++ { // unlike java brackets are not required for for loop but braces are mandatory
// 		// repeated = repeated + character // assign value to variable
// 		repeated += character // append value to variable
// 	}
// 	return repeated
// }

// comment out the custom implementation above in favour of standard library
// drops brnchmark from 80->30ns!
func Repeat(character string, iterations int) string {
	// use string library to repeat the character n times
	return strings.Repeat(character, iterations)
}
