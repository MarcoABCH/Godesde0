package teclado

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros(){
	
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese numero 1: ")
	if scanner.Scan() {
		//Nota: al asignar se pone := cuando las variables no existen previamente
		//numero1, err := strconv.Atoi(scanner.Text())
		//Como numero1 y err ya fueron declarados se puede utilizar el = al asignar
		numero1, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("Se aborto la aplicacion, ya que el dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese número 2: ")
	if scanner.Scan() {		
		numero2, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("Se aborto la aplicacion, ya que el dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda: ")
	if scanner.Scan() {		
		leyenda= scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)
}