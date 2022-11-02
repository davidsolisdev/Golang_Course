package variables

import "fmt"

func Vars(){
  
  // Declaracion simple de variables 
  var nombre string = "David"

  // Declaracion de variable en bloque
  var (
    apellido string = "Solis"
    edad uint8 = 22
  )

  // Declaracion de variables forma corta o dinamicamente
  userAlias := "davidsolisdev"

  // Uso de las variables para evitar errores
  fmt.Println(nombre, apellido, userAlias, edad)
  
}