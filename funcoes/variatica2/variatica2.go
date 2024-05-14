package main()
//slice
func imprimirAprovados(aprovados ...string){
	fmt.Println("Lista de Aprovados:  ")
	for i, aprovado := range aprovados{
		fmt.Printf("%d, %s", i+1, aprovado)
	}
}

func main(){
	aprovados := []string{"Maria","Claudia", }
}