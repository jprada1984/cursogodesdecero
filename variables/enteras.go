package variables

import "fmt"

func MostrarEntero() {
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)
	intde64 = int64(intcomun)

	fmt.Println("intcomun=", intcomun)
	fmt.Println("intde32=", intde32)
	fmt.Println("indde64=", intde64)
}
