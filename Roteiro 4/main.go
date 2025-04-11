package main

import (
	"fmt"
)

func main() {
	fmt.Println("G3T L0W")

	// Uma pilha (stack) é uma estrutura de dados do tipo LIFO (Last In, First Out), ou seja, o
	//último elemento inserido é o primeiro a ser removido

	//A pilha suporta as seguintes operações básicas:
	//	● Push: Adiciona um elemento ao topo da pilha.
	//	● Pop: Remove e retorna o elemento do topo da pilha.
	//	● Top (Peek): Retorna o elemento do topo sem removê-lo.
	//	● IsEmpty: Verifica se a pilha está vazia

	s := Stack{}

	s.Push(10)
	s.Push(69)
	s.Push(96)

	top, exists := s.Top()
	if exists {
		fmt.Println("Elemento no topo: ", top)
	} else {
		fmt.Println("A Lista está vazia: ")
	}

	for !s.Isempty() {
		val, _ := s.Pop()
		fmt.Println("Removido: ", val)
	}

}

// Vamos implementar uma pilha usando uma estrutura (struct) simples. Utilizaremos um
// slice para armazenar os elementos.
type Stack struct {
	items []int // defininido array de itens
}

// Push para adicionar um elemento ao topo da pilha
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop para Remover e retornar o elemento do topo da pilha
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false // retorna falso se a pilha estiver vazia
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex] // removação ultimo elemento
	return item, true
}

// Top para retornar o elemento do topo sem remover
func (s *Stack) Top() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

// Isempty verifica se a lista esta vazia
func (s *Stack) Isempty() bool {
	return len(s.items) == 0
}
