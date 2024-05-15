package iteraciones

import (
	"fmt"
)

func Iterar(){
	//El for en Go es muy versatil se puede usar asi, version larga
	/*i:=0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	*/
	//version mas corta
	/*for i:=0; i < 10; i++ {
		fmt.Println(i)		
	}
	*/
	/*
	for i:=0; i < 100; i+=5 {//que salte de 5 en 5
		fmt.Println(i)		
	}
	*/
	for i:=10; i > 0; i-- {//que salte de 5 en 5
		//continue permite que se salte este numero y no lo muestra, quiere decir "continua al que sigue y no hagas nada mas"
		if i==6 {
			continue
		}
		fmt.Println(i)		
	}
}