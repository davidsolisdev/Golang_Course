package conversiondedatos

func StringAndInt(){
  // Convertir de string a int
	numero, err := strconv.Atoi("1")

	// Convertir de int a string
	texto := strconv.Itoa(1)

	fmt.Println(numero, texto, err.Error())
}