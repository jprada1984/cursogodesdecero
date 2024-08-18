package main

import (
	"fmt"

	"github.com/jprada1984/cursogodesdecero/variables"
)

func main() {
	// variables.MostrarEntero()
	correcto, valor := variables.ConviertoaTexto(150)
	fmt.Println(correcto)
	fmt.Println(valor)
}
