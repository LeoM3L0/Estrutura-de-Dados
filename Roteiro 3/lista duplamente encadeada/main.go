package main

import (
	"fmt"
)

func main() {
	fmt.Println("H3LL0 MY G0D")

	//Uma Lista Duplamente Encadeada é uma variação onde cada nó contém:
	//1. Um valor (data)
	//2. Um ponteiro para o próximo nó (next)
	//3. Um ponteiro para o nó anterior (prev)

	list := DoublyLinkedList{}
	list.Append(69)
	list.Append(96)
	list.Append(169)
	list.Append(8960)

	fmt.Println("Lista do inicio ao fim")
	list.PrintForward()

	fmt.Println("\n Lista do fim para o inicio")
	list.PrintBackward()

	list.RemoveLast()
	fmt.Println("\n lista após remover o ultimo elemento")
	list.PrintForward()

	// Criamos a DoublyLinkedList, que permite percorrer nos dois sentidos.
	//● Implementamos funções para:
	//○ Adicionar elementos (Append)
	//○ Remover o último elemento (RemoveLast)
	//○ Percorrer a lista em ambos os sentidos (PrintForward e
	//PrintBackward)
	//Agora podemos navegar tanto para frente quanto para trás!
}

// Implementação da Lista Duplamente Encadeada

type Node struct {
	data int
	next *Node
	prev *Node
}

// Estrutura da lista duplamente encadeada
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// Adiciona um elemento no final da lista
func (l *DoublyLinkedList) Append(data int) {
	newNode := &Node{data: data}
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	l.tail.next = newNode
	newNode.prev = l.tail
	l.tail = newNode
}

// Remove o último elemento da lista
func (l *DoublyLinkedList) RemoveLast() {
	if l.tail == nil {
		return
	}
	if l.tail == l.head { // Apenas um elemento na lista
		l.head = nil
		l.tail = nil
		return
	}
	l.tail = l.tail.prev
	l.tail.next = nil
}

// Exibe os elementos da lista do início ao fim
func (l *DoublyLinkedList) PrintForward() {
	current := l.head
	for current != nil {
		fmt.Print(current.data, " <-> ")
		current = current.next
	}
	fmt.Println("nil")
}

// Exibe os elementos da lista do fim para o início
func (l *DoublyLinkedList) PrintBackward() {
	current := l.tail
	for current != nil {
		fmt.Print(current.data, " <-> ")
		current = current.prev
	}
	fmt.Println("nil")
}
