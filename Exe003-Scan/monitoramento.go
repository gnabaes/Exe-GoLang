package main

import "fmt"

func main() {
	fmt.Println("1 - Iniciar Monitoramento ")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	var comando int
	fmt.Scanf("%d", &comando)                                         // em lugar de usar a Scanf pode ser usada a função Scan (&comando) sem especificar.
	fmt.Println(" o Endereço da minha variavel comando e ", &comando) // quando colocar &Comando esta indicando que vai ser um ponteiro para armazenar na memoria
	fmt.Println("O comando escolhido foi", comando)
}
