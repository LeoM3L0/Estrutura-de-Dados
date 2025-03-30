package main

import (
	"fmt"
)

func main() {
	fmt.Println("HELL0 W0RLD")

	//Atribuindo os números de meu R.A em um vetor de forma manual
	var array [6]int

	array[0] = 1
	array[1] = 2
	array[2] = 4
	array[3] = 1
	array[4] = 7
	array[5] = 1

	//Acessando valores de forma manual
	fmt.Printf("%d", array[0])
	fmt.Printf("%d", array[1])
	fmt.Printf("%d", array[2])
	fmt.Printf("%d", array[3])
	fmt.Printf("%d", array[4])
	fmt.Printf("%d", array[5])
	fmt.Printf("\n\n")

	// Acessando valor atraves de for
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d", array[i])
	}

	fmt.Printf("\n\n")

	// Acessando Array atráves de range
	for _, i := range array {
		fmt.Printf("%d", i)
	}

	// DB Alunos
	var alunos [3][3]string

	for i := 0; i < len(alunos); i++ {
		fmt.Printf("\n\nInsira o Índice do Aluno: %d\n", i+1)
		fmt.Scan(&alunos[i][0])

		fmt.Printf("Insira o RA do Aluno: %d\n", i+1)
		fmt.Scan(&alunos[i][1])

		fmt.Printf("Insira o Nome do Aluno: %d\n", i+1)
		fmt.Scan(&alunos[i][2])
	}

	fmt.Println("Dados dos Alunos:")
	fmt.Println("========================")
	fmt.Printf("%-8s %-8s %-8s \n", "Índice", "RA", "NOme")

	for i := 0; i < len(alunos); i++ {
		fmt.Printf("%-8s %-8s %-8s \n", alunos[i][0], alunos[i][1], alunos[i][2])
	}

	// Structs
	Produto1 := Produtos{
		Nome:       "Ajax",
		Preço:      69.69,
		Quantidade: 69,
	}
	Produto2 := Produtos{
		Nome:       "Rothmans Red",
		Preço:      7,
		Quantidade: 45,
	}
	Produto3 := Produtos{
		Nome:       "Lucky Strike Red",
		Preço:      7,
		Quantidade: 32,
	}

	listaprodutos := []Produtos{Produto1, Produto2, Produto3}
	fmt.Println("Lista de Produtos")

	for i, p := range listaprodutos {
		fmt.Printf("\nProduto %d: \n", i+1)
		fmt.Println("Nome:", p.Nome)
		fmt.Println("Preço:", p.Preço)
		fmt.Println("Quatidade", p.Quantidade)
	}

}

type Produtos struct {
	Nome       string
	Preço      float64
	Quantidade int
}
