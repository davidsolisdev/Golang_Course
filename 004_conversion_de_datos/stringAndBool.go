package conversiondedatos

import (
	"fmt"
	"strconv"
)

func StringAndBool() {
	// * convert string to bool
	// 0, 1, t, f, T, F, true, false, TRUE, FALSE, True, False
	falso, err := strconv.ParseBool("false")

	// * convert bool to string
	verdadero := strconv.FormatBool(true)

	fmt.Println(falso, verdadero, err)
}
