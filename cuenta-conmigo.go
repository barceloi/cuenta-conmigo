package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	letterA = `
  A
 A A
AAAAA
A   A
A   A`
	letterB = `
 	Braquiosaurio

	  _
 .~q \,
{__,  \
    \' \
     \  \
      \  \
       \  \._            __.__
        \    ~-._  _.==~~     ~~--.._
         \        '                  ~-.
          \      _-   -_                \.
           \    /       }        .-    .  \
            \. |      /  }      (       ;  \
              \|     /  /       (       :    \
               \    |  /        |      /       \    =
                |     /\-.______.\     |^-.      \
                |   |/           (     |    .      \_
                |   ||            ~\   \      '._    \-.._____..----.._=__
                |   |/             _\   \      =~-.__________.-~~~~~~~~~'''
              .o'___/            .o______}`
	letterT = `
	Triceratops
                       _. - ~ ~ ~ - .
   ..       __.    .-~               ~-.
   ((\     /   \}.~                     \.
    \\\   {     }               /     \   \
(\   \\~~^      }              |       }   \
 \\.-~ -@~      }  ,-.         |       )    \
 (___     ) _}   (    :        |    / /      \.
  \----._-~.     _\ \ |_       \   / /- _      \.
         ~~----~~  \ \| ~~--~~~(  + /     ~-.     ~- _
                   /  /         \  \          ~- . _ _~_-_.
                __/  /          _\  )
              .<___.'         .<___/`
	letterTOption2 = `
	Triceratops 2
  TT       TT
  TT       TT
   TTTTTTTTT
   TTTTTTTTT
        TTTT
        TTTT
`
	letterTOption3 = `
	Triceratops 3
  TTTTTTTTTTTTT
  TTTTTTTTTTTTT
        TTTT
        TTTT
        TTTT
        TTTT
`
)

func main() {
	fmt.Println("Había una vez, en un lugar muy muy lejano. Un...")

	var inputLetter string

	asciiOptions := map[string][]string{
		"A": {letterA, letterA}, // Incluir ambas opciones para "A" (mayúscula y minúscula)
		"a": {letterA, letterA}, // Incluir ambas opciones para "a" (mayúscula y minúscula)
		"B": {letterB},
		"b": {letterB},
		"T": {letterT, letterTOption2, letterTOption3}, // Agregar más opciones para la letra "T"
		"t": {letterT, letterTOption2, letterTOption3}, // Agregar más opciones para la letra "t"
	}

	seed := time.Now().UnixNano()

	r := rand.New(rand.NewSource(seed))

	for {
		fmt.Print("Ingresa una letra o 'fin' para terminar: ")
		fmt.Scanln(&inputLetter)

		if inputLetter == "fin" {
			break
		}

		options, found := asciiOptions[inputLetter]

		if !found {
			fmt.Println("La letra ingresada no tiene una imagen ASCII asociada.")
			continue
		}

		selectedASCII := options[r.Intn(len(options))]
		fmt.Println(selectedASCII)

		fmt.Println("Ingresa otra letra para ver con qué, quién o qué lugar se encuentra nuestro personaje (o escribe 'fin' para terminar): ")

	}

	fmt.Println("¡Hasta luego!")

}
