package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //o compilador vai contar

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}
	for _, num := range numeros {
		fmt.Println(num)
	}
}
