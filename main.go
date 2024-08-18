package main

import (
	"fmt"

	"github.com/jprada1984/cursogodesdecero/ejercicios"
)

func main() {
	//correcto, valor := variables.ConviertoaTexto(150)
	//fmt.Println(correcto)
	// fmt.Println(valor)

	// if os := runtime.GOOS; os == "Linux." || os == "darwin" {
	// 	fmt.Println("Esto no es Windows, es ", os)
	// } else {
	// 	fmt.Println("Esto es Windows")
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	numero, respuesta := ejercicios.Funcion01("1011")
	fmt.Println("el parametro es:", numero)
	fmt.Println("la respuesta es:", respuesta)
}
