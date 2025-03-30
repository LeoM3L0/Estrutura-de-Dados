package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	fmt.Printf("\n\n")

	// Declaração e  Inicialização de Arrays
	var array [5]int
	fmt.Println(array)

	fmt.Printf("\n\n")

	arr := [3]int{10, 25, 50}
	fmt.Println(arr)

	fmt.Printf("\n\n")

	// Acessando com Valores
	arr[1] = 9
	fmt.Println(arr[1])
	fmt.Printf("\n\n")

	// Percorrendo Arrays com for Tradiconal
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Printf("\n\n")

	//Range é usado para percorrer arrays, slices, maps, canais e strings
	// De maneirea efienente e legivel retornando 2 valores: Índice ou Chave e Valor Armazenado
	for i, v := range arr {
		fmt.Printf("Índice: %d, Valor: %d\n", i, v)
	}
	fmt.Printf("\n\n")

	// OBS: Sempre que houver mais de um retorno, podemos otimizar utilizando _ (under score)
	for _, v := range arr {
		fmt.Printf("VAlor: %d\n", v)
	}
	fmt.Printf("\n\n")

	//Declarando e inicializando MAtriz
	var matriz [2][3]int
	fmt.Println(matriz)

	// Inicializando com Valores
	matriz2 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(matriz2)

	// Acessando e modificando valores
	matriz3 := [2][2]int{{6, 9}, {9, 6}}
	fmt.Println(matriz3[1][0])

	matriz3[0][1] = 14
	fmt.Println(matriz3)

	// Percorrendo Matriz
	for i := 0; i < len(matriz3); i++ {
		for j := 0; j < len(matriz3[i]); j++ {
			fmt.Printf("%d ", matriz3[i][j])
		}
		fmt.Println()
	}

	for i, linha := range matriz3 {
		for j, valor := range linha {
			fmt.Printf("Matriz:[%d][%d] = %d \n", i, j, valor)
		}
	}

	// Struct se Assemelha a Classes e outras linguagens
	type Pessoa struct {
		Nome  string
		Idade int
	}

	p := Pessoa{
		Nome:  "jao",
		Idade: 69,
	}

	fmt.Println(" Meu Nome é:", p.Nome, "Tenho", p.Idade, "Anus")

}
