package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
)

func main() {

	defer fmt.Println("world")

	var c, python, java = true, false, "no!"

	var i, j int = 1, 2
	k := 3

	fmt.Println(" Hello World")
	fmt.Println("Meu numero favorito é: ", rand.Intn(1000))
	fmt.Println(math.Pi)
	fmt.Println(adicao(50, 30))

	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println(split(20))

	fmt.Println(i, j, k, c, python, java)

	sum := 0

	for l := 0; l < 10; l++ {
		sum += l
	}
	fmt.Println(sum)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Print("GO runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X, ")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}

	fmt.Println("hello")
}

func adicao(x int, y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
