package conversiondedatos

import (
	"fmt"
	"strconv"
)

func StringAndFloat() {
	var numero64String string = "786543.12345678"
	var numero32String string = "54374.9786"

	// * convert string to float64
	nFloat64, err := strconv.ParseFloat(numero64String, 64)
	fmt.Printf("El resultado es: %v , %T \n", nFloat64, nFloat64) // 786543.12345678 , float64

	// * convert string to float32
	n32, err := strconv.ParseFloat(numero32String, 32)
	var nFloat32 float32 = float32(n32)
	fmt.Printf("El resultado es: %v , %T \n", nFloat32, nFloat32) // 54374.98 , float32

	if err != nil {
		panic(err.Error())
	}
}
