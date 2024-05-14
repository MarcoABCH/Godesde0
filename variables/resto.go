package variables

import (
	"fmt"
	"time"
	"strconv"
)

var Nombre string//Es publica para todo el proyecto donde se importe
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables (){
	Nombre = "Marco"
	Estado = true
	Sueldo = 40000.00
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println( Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoaTexto(numero int) (bool,string) {
	texto := strconv.Itoa(numero)
	return true,texto
}
