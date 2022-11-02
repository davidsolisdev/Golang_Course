package conversiondedatos

import (
	"fmt"
	"strconv"
)

func StringAndInt() {
	// * convert string to int
	numero, err := strconv.Atoi("1")
	numeroo, err := strconv.ParseInt("32767", 10, 16)

	// * convert int to string
	texto := strconv.Itoa(1)
	intTexto := strconv.FormatInt(25, 10)

	fmt.Println(numero, numeroo, texto, intTexto, err)
}
