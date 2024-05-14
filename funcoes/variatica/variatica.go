package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))

}

func main() {
	fmt.Printf("Media: %f", media(7.8, 8.1, 9.9, 8.5, 9.8))
}
