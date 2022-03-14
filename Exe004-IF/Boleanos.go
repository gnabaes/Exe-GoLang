package main

import "fmt"

func main() {
	fmt.Println("1 - Iniciar Monitoramento ")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	var comando int
	fmt.Scanf("%d", &comando) // em lugar de usar a Scanf pode ser usada a função Scan (&comando) sem especificar.
	fmt.Println("O comando escolhido foi", comando)

	if comando == 1 {
		fmt.Println("Monitorando ....")
	} else if comando == 2 {
		fmt.Println("Exibindo logs....")
	} else if comando == 0 {
		fmt.Println("Sáindo do programa ....")
	} else {
		fmt.Println(" ERRO não conheço este comando digitado")
	}

}
