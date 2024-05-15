package files

import (
	"github.com/MarcoABCH/Godesde0/ejercicios"
	"os"	
	"fmt"	
	"bufio"
)

var fileName string = "./files/txt/tabla.txt" 

func GrabaTabla(){
	var texto string = ejercicios.TablasdeMultiplicar()
	//Vamos a mandar la tabla de multiplicar anterior a un archivos

	archivo, err:= os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)

	archivo.Close()//siempre debe de borrarse un archivo creado
	
}


func SumaTabla(){
	var texto string = ejercicios.TablasdeMultiplicar()

	if !Append(texto) {
		fmt.Println("Error al concatenar el contenido ")		
	}	
}

func Append(texto string) bool {
	arch, err:= os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append")
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}

	arch.Close()

	return true
}

func LeoArchivo(){
	archivo, err:= os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(">" + registro)
	}

	//Siempre hay que hacer el Close en el manejo de archivos p liberar memoria
	archivo.Close()
}