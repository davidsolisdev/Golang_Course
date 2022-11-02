package conversiondedatos

import "fmt"

func Integers(){
  var edad int64 = 22
	var age int8
	var ageUnsigned uint8

	age = int8(edad)

	ageUnsigned = uint8(edad)

	fmt.Println(age, ageUnsigned)
}