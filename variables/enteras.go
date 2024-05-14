package variables

import "fmt"

//La primera letra mayuscula indica que es una funcion publica
func MuestroEnteros() {//Aqui en Go siempre la 1er llave va arriba no se baja
	//Las variables se declaran por asignacion
	var intcomun int //Al declararla por default intcomun vale 0, no le asigna nullos el Go

	//Estas variables se declaran al momento de asignar, va tomar el tipo de dato del valor que se le asigne
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)

}

//aqui se le pone package seguida de la carpeta donde esta nuestro archivo enteras.go
