package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Println("Informe um numero: ")
	fmt.Scan(&x)

	fmt.Println(ParouImpar(x))

	var y int
	fmt.Println("Insira um número Que somara todos os anteriores até ele: ")
	fmt.Scan(&y)
	fmt.Println(SomaAte(y))

	fmt.Println("Insira nota de 1 a 10: ")
	fmt.Scan(&x)
	fmt.Println(ClassificarNota(x))

}

func ParouImpar(n int) string {
	if n%2 == 0 {
		return "Par"
	} else {
		return "Impar"
	}
}

func SomaAte(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	return sum
}

func ClassificarNota(nota int) string {
	switch {
	case nota >= 9 && nota <= 10:
		return "Excelente"
	case nota >= 7 && nota <= 8:
		return "Bom"
	case nota >= 5 && nota <= 6:
		return "Regular"
	case nota >= 3 && nota <= 4:
		return "Ruim"
	case nota >= 0 && nota <= 3:
		return "Muito Ruim"
	default:
		return "Nota de 1 a 10"
	}
}
