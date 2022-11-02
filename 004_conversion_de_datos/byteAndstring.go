package conversiondedatos

import "fmt"

func ByteAndString(){
  var nombre string = "David"
  var name []byte

  name = []byte(nombre)

  nombre = string(name)

  fmt.Println(nombre, name)
}