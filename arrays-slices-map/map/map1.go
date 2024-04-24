package main

import "fmt"

func main1() {
	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[223456789] = "Luiz"
	aprovados[323456789] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf(" %s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[323456789])
}
