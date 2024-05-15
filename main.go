package main

//palabra d einicio siempre es package seguido del archivo actual en este caso main
import (
	"fmt"	
	"github.com/MarcoABCH/Godesde0/ejercicios"
)

func main(){
	/*
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado, texto)
	*/

	numero1, string1 := ejercicios.EsMayorMenor("90")
	fmt.Println(numero1, string1)

	/*
	if os:= runtime.GOOS; os == "linux" || os == "OS X." {//Pregunta y asigna a la vez, eso es menos código
		fmt.Println("No eres Windows, que SO eres?")
	} else {
		fmt.Println("Hola Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Hola Linux")
	case "darwin":
		fmt.Println("Hola Darwin")
	case "os x":
		fmt.Println("Hola os X")
	default:	
		fmt.Printf("%s \n", os)
	}
	*/
		
}