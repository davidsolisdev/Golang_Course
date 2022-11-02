package conversiondedatos

import "fmt"

func ByteAndString() {
	var nombre string = "David"
	var name []byte

	// * convert string to []byte
	name = []byte(nombre)

	// * convert []byte to string
	nombre = string(name)

	fmt.Println(nombre, name)
}
