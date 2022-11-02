package constantes

import "fmt"

func Constants(){
  
  // Declaracion normal de constantes
  const nombre = "David"

  // Declaracion en bloque de constantes
  const (
    apellido string = "Solis"
    edad uint8 = 22
  )

  // Las constantes no es necesario utilizarlas
  fmt.Println("Mi nombre es:", nombre, apellido, "mi edad es:", edad)
  
}