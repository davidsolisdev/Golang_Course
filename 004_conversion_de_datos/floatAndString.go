package conversiondedatos

import (
	"fmt"
	"strconv"
)

func FloatAndString() {
	var numero64 float64 = 83456435.2345567
	var numero32 float32 = 534454.974353

	// * convert float64 to string
	n64String := strconv.FormatFloat(numero64, 'f', 4, 64)
	fmt.Printf("El resultado es: %v , %T \n", n64String, n64String) // 83456435.2346 , string

	// * convert float32 to string
	n32String := strconv.FormatFloat(float64(numero32), 'f', 4, 32)
	fmt.Printf("El resultado es: %v , %T \n", n32String, n32String) // 534455.0000 , string
}
