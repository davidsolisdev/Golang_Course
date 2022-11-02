package conversiondedatos

import "fmt"

func Integers() {
	var edad int64 = 22
	var age int8
	var ageUnsigned uint8

	// * convert int64 to int8
	age = int8(edad)

	// * convert int64 to uint8
	ageUnsigned = uint8(edad)

	fmt.Println(age, ageUnsigned)
}
