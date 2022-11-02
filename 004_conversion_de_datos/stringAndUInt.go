package conversiondedatos

import (
	"fmt"
	"strconv"
)

func StringAndUInt() {
	// * convert string to uint
	numero, err := strconv.ParseUint("32767", 10, 16)

	// * convert uint to string
	uintTexto := strconv.FormatUint(25, 10)

	fmt.Println(numero, uintTexto, err)
}
