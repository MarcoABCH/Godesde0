package ejercicios

import "fmt"
import "strconv"

func EsMayorMenor(valor string) (int, string){
	numero, err := strconv.Atoi(valor)

	if err == nil {
		if numero > 100 {
			fmt.Println("Es mayor a 100")
		} else {
			if numero < 100 {
				fmt.Println("Es menor a 100")
			}else{
				fmt.Println("No es un valor valido para la prueba, solo se admiten menor a 100 o mayor a 100")
			}
		}
	}else {
		fmt.Println("Error al convertir el valor a entero")
	}

	return numero, valor
}