package main

import (
	"fmt"
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
)

func main() {
	fmt.Println("Había una vez, en un lugar muy muy lejano. Un...")

	var inputLetter string

	for {
		fmt.Print("Ingresa una letra o 'fin' para terminar: ")
		fmt.Scanln(&inputLetter)

		if inputLetter == "fin" {
			break
		}

		switch inputLetter {
		case "A", "a":
			fmt.Println(letterA)
		case "B", "b":
			fmt.Println(letterB)
		case "T", "t":
			fmt.Println(letterT)

		default:
			fmt.Println("La letra ingresada no tiene una imagen ASCII asociada.")
		}

		fmt.Println("Ingresa otra letra para ver con qué, quién o qué lugar se encuentra nuestro personaje (o escribe 'fin' para terminar): ")

	}

	fmt.Println("¡Hasta luego!")

}
