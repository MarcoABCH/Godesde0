package main

//palabra d einicio siempre es package seguido del archivo actual en este caso main
import (
	"fmt"
	"github.com/MarcoABCH/Godesde0/variables"
)

func main(){

	estado, texto := variables.ConviertoaTexto(1588)

	fmt.Println(estado, texto)
}