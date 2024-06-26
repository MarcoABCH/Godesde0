package ejercicios

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

var tabla int
var err error
var texto string

func TablasdeMultiplicar() string {//Regresa un string como e sun solo parametro no se pone entre ()
	scanner := bufio.NewScanner(os.Stdin)
	
	//Mi solucion
	for i:=0; i<2; i++ {
		fmt.Println("Que tabla de multiplicar deseas ?: ")

		if scanner.Scan() {
			tabla, err = strconv.Atoi(scanner.Text())

			if err != nil {//Si trae algun error
				i=0
				fmt.Println("Por favor capture un numero entre 1 al 10")
				continue
			}else {
				if tabla<0 || tabla>10 {
					i=0
					fmt.Println("Por favor capture un numero entre 1 al 10")
					continue
				}
				break
			}
		}
	}

	//solucion instructor
	/*for {
		fmt.Println("Que tabla de multiplicar deseas ?: ")
		if scanner.Scan() {
			tabla, err = strconv.Atoi(scanner.Text())

			if err != nil {//Si trae algun error								
				continue
			}else {				
				break								
			}
		}
	}
	*/

	

	if tabla>0 && tabla<=10 {
		fmt.Println("La tabla del " + strconv.Itoa(tabla) + " es: ")
		
		for k:=1; k<=10;k++ {
			//Mis solucion
			//fmt.Println(strconv.Itoa(tabla) + " X " + strconv.Itoa(k) + " = " + strconv.Itoa(tabla * k))

			//solucion instructor
			texto += fmt.Sprintf("%d X %d = %d \n", tabla,k,tabla*k)
		}		
	}
	return texto
}