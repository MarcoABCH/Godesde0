package main

//palabra d einicio siempre es package seguido del archivo actual en este caso main
import (		
	"github.com/MarcoABCH/Godesde0/files"	
)

func main(){
	/*
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado, texto)
	*/

	//numero1, string1 := ejercicios.EsMayorMenor("90")
	//fmt.Println(numero1, string1)

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
		
	//teclado.IngresoNumeros()
    /*tabla := 1
	for {
		
		if tabla == 11 {						
			break
		}else {			
			fmt.Println("Estoy en la tabla " + strconv.Itoa(tabla))	
		}
		
		tabla = tabla + 1
	}*/

	//iteraciones.Iterar()
	//fmt.Println(ejercicios.TablasdeMultiplicar())
	//files.GrabaTabla()

	//files.SumaTabla()
	files.LeoArchivo()
}