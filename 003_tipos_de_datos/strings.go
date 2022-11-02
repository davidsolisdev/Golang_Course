package tiposdedatos

import "fmt"

func Strings(){
  
  // Variables de tipo string
  const nombre string = "David"

  const apellido string = "Solis"

  userAlias := "davidsolisdev"

  const textoLiteral string = `En este texto se respeta el salto de linea
			y se muestra de la misma manera en la que se codificó`

  // Uso de las variables para compilar sin errores
  fmt.Println(userAlias)
  
}