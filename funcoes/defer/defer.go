package main

import "fmt"

// atrasa uma determinada função, fecha algum recurso que voce abriu
// atrasa algo para o ultimo momento possivel dentro do metodo
func obterValorAprovado(numero int) int {
	defer fmt.Println("----- fim!!! ------")
	if numero >= 5000 {
		fmt.Println("Valor alto...")
		return 5000
	} else {
		fmt.Println("Valor baixo...")
		return numero
	}
}
func obterValorAprovado2(numero int) int {
	if numero >= 5000 {
		fmt.Println("Valor alto...")
		return 5000
	} else {
		fmt.Println("Valor baixo...")
		return numero
	}
}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
	fmt.Println("------------------------ ")
	fmt.Println(obterValorAprovado2(6000))
	fmt.Println(obterValorAprovado2(3000))
}
