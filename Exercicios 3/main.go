package main

import (
	"fmt"
)

func main() {
	fmt.Println("H3LL0 MY L0RD")

	var b float64 = 31.2121
	var p *float64 = &b

	fmt.Println("valor de b antes de modificiado: ", b)

	*p = 123.4321

	fmt.Println("Valor de b modificado por ponteiro", b)

	z := 69
	c := 96

	fmt.Println("Valores de z e c antes de Swap: ", z, c)
	swap(&z, &c)
	fmt.Println("Valores de z e c depois de Swap: ", z, c)
}

func swap(num1, num2 *int) {
	tmp := *num1
	*num1 = *num2
	*num2 = tmp
}
