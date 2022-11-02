package tiposdedatos

func CompTiposDatos(){
  userName := "davidsolisdev"

	var tipo reflect.Kind = reflect.TypeOf(userName).Kind()
	
	if tipo == reflect.String {
		fmt.Println("si es string")
	}
  
	fmt.Println(tipo) // string
}