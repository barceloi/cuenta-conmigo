package main

// import(
// 	"testing"
// )

// func TestMain(t *testing.T) {
// 	result := main( inputLetter: z)
// 	if result != ``
// }

import (
	"os"
	"strings"
	"testing"
)

func TestHandleUserInput(t *testing.T) {
	// Creamos un buffer para capturar la salida de la función
	var output strings.Builder
	oldOutput := output
	output = strings.Builder{}
	defer func() { output = oldOutput }()

	tests := []struct {
		input      string
		shouldExit bool
	}{
		{input: "A\n", shouldExit: false},
		{input: "X\n", shouldExit: false},
		{input: "z\n", shouldExit: false},
		{input: "Ñ\n", shouldExit: false},
		{input: "fin\n", shouldExit: true},
		{input: "end\n", shouldExit: true},
		{input: "Y\nfin\n", shouldExit: true},
	}

	for _, test := range tests {
		inputReader := strings.NewReader(test.input)
		oldInput := inputReader
		fmtReader := strings.NewReader("random")
		oldFmt := fmtReader
		defer func() { inputReader, fmtReader = oldInput, oldFmt }()

		// Reemplazamos las entradas y salidas por nuestros readers y buffer
		inputReader, os.Stdin = os.Stdin, inputReader
		fmtReader, os.Stdout = os.Stdout, fmtReader
		handleUserInput()
		// Comprobamos si la salida es la esperada
		actualOutput := output.String()
		if test.shouldExit {
			expectedOutput := "Ingresa una letra o 'fin' para terminar: \n / Enter a letter or type 'end' to finish: ¡Hasta luego! \n Good bye!\n"
			if actualOutput != expectedOutput {
				t.Errorf("Para la entrada %q, se esperaba la salida %q, pero obtuvimos %q", test.input, expectedOutput, actualOutput)
			}
		} else {
			// La función no debe haber salido, así que esperamos un mensaje adicional
			if !strings.Contains(actualOutput, "Ingresa otra letra") {
				t.Errorf("Para la entrada %q, se esperaba un mensaje adicional, pero obtuvimos %q", test.input, actualOutput)
			}
		}

		output.Reset() // Limpiamos el buffer para el siguiente test
	}
}
