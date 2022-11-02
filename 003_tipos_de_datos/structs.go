package tiposdedatos

import "fmt"

func Structs(){
  
  type Persona struct {
    Nombre string
    Apellido string
    Edad uint8
  }
  
  var David Persona = Persona{
    Nombre: "David",
    Apellido: "Solis",
    Edad: 22,
  }

  fmt.Println(David)
}