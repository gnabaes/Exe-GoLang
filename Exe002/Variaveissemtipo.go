package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome = "Antonella" // podemos fazer comentarios igual a javascript no VSC.
	var idade = 7          // podemos retirar o var e utilziando :=  idade := no go lang
	cor := "branca"
	fmt.Println("o nome da minha filha", nome, " ela tem idade de  ", idade, " e tem pele de cor ", cor)
	fmt.Println("o tipo da variavel nome é  ", reflect.TypeOf(nome), " o tipo da variavel idade é de tipo", reflect.TypeOf(idade))

}
