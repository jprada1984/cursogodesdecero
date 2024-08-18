package ejercicios

import (
	"strconv"
)

func Funcion01(param string) (int, string) {
	num, error := strconv.Atoi(param)

	if error != nil {
		return 0, "Hubo un error " + error.Error()
	}
	if num > 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menos a 100"
	}

}
