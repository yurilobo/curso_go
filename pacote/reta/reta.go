package main

import "math"

//inicializando com letra maiuscula é visivel dentro e foa do pacote
//se inicializar com letra minuscula, o pacote é privado dentro do pacote, nao dentro do arquivo
//ponto ou _Ponto - gerara algo privado

// ponto possui uma coordenada  do plano cartesiana
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// distancia e responsavel por calculo a distancia entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
