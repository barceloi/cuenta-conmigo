package main

import (
	"strings"
	"testing"
)

// TestHandleUserInput tests the handleUserInput function.
func TestHandleUserInput(t *testing.T) {
	// Create a strings.Reader to read user input from the console.
	inputReader := strings.NewReader("A\nend\n")

	// Create a strings.Builder to collect the output.
	outputBuilder := &strings.Builder{}

	// Call handleUserInput and pass the inputReader and outputBuilder as arguments.
	handleUserInput(inputReader, outputBuilder)

	// Define the expected output for the given input.
	expectedOutput := "Ingresa una letra o 'fin' para terminar: \n / Enter a letter or type 'end' to finish: \n" +
		"Había una vez, en un lugar muy muy lejano. Había un ... / Once upon a time, in a land far far away. There was a ...\n" +
		"Ingresa otra letra para ver con qué, quién o qué lugar se encuentra nuestro personaje (o escribe 'fin' para terminar): \n / Enter another letter to see with what, who, or what place our character encounters (or type 'end' to finish): \n" +
		"¡Hasta luego! \n Good bye!\n"

	// Compare the actual output with the expected output.
	if outputBuilder.String() != expectedOutput {
		t.Errorf("handleUserInput() produced incorrect output. \n\nExpected:\n%s\n\nActual:\n%s", expectedOutput, outputBuilder.String())
	}
}
