package metodosparamaps

import "fmt"

func ListarPropiedades() {

	var data map[string]interface{} = map[string]interface{}{
		"id":   1,
		"name": "admin",
	}

	for key, value := range data {
		fmt.Println(key, value)
	}

}
