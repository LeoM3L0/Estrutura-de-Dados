package main

import (
	"fmt"
)

func main() {
	fmt.Println("y0u r34dy f0r th1s")

	words := []string{"radar", "arara", "golang", "level", "hello", "paralelepipedo", "mascarenhas"}

	for _, word := range words {
		if IsPalindromo(word) {
			fmt.Printf("\"%s\" é um palindromo\n", word)
		} else {
			fmt.Printf("\"%s\" não é palindromo\n", word)
		}
	}
}

// definindo a pilha

type Stack struct {
	items []rune
}

// Push adiciona caractere a ppilha
func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

// pop remove e retorna o caractere do top oda pilha
func (s *Stack) Pop() (rune, bool) {
	if len(s.items) == 0 {
		return ' ', false
	}
	// pega o ultimo elemento
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

// função para verificar se uma palavra é palindromo usando pilha
func IsPalindromo(word string) bool {
	stack := Stack{}

	// empliha todos os caracteres de palavras
	for _, char := range word {
		stack.Push(char)
	}

	// desempilha para construir a palavra invertida
	reversed := ""
	for range word {
		char, ok := stack.Pop()
		if ok {
			reversed += string(char)
		}
	}
	// compara a palavra original com a invertida
	return word == reversed
}
