package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel Silva": 15456.78,
			"Guga Pereira":  8546.58,
		},
		"J": {
			"Joao Joao Silva": 16356.78,
			"Julio Guga":      4546.58,
		},
		"P": {
			"Pedro Paulo": 8456.78,
		},
	}
	delete(funcsPorLetra, "J")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
