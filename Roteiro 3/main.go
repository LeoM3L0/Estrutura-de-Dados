package main

import (
	"fmt"
)

func main() {
	fmt.Println("H3LL0 W0RLD")

	// Ponteiros são variáveis que armazenam endereço de memoria de outra variavel
	// declarando ponteiro
	// Operador & retorna o enderenço da memoria de uma variavel
	// Operador * acessa o valor armazenado no endereço

	// Passagem por valor: uma copia do valor da variavel é passada para a função
	// Passagem por referência: o endereço de memoria da variável é passado permitindo que função modifique o valor original

	var a int = 42
	var p *int = &a

	fmt.Println("Valor de a: ", a)
	fmt.Println("Endereço de a:", &a)
	fmt.Println("Valor de p (endereço de a): ", p)
	fmt.Println("Valor a pontado por p: ", *p)

	// Modificando Valores através de Ponteiros
	var b float64 = 69.69
	var p1 *float64 = &b

	fmt.Println("vlor de a antes de modificado", *p1)
	*p1 = 96.96
	fmt.Println("Valor de a modificado por ponteiro", *p1)

	// PONTEIROS EM FUÇÕES
	x := 10
	increment(x)
	fmt.Println("Fora da func increment: ", x)

	y := 68
	incrementobyponteiro(&y)
	fmt.Println("fora da func incrementobyponteiro: ", y)

	// LISTAS ENCADEADAS
	// LISTA ENCADEADAS É UMA ESTRUTURA DE DADOS ONDE CADE ELEMENTO (NÓ) CONTEM
	// UM VALOR (DATA)
	// UM PONTEIRO PARA O PROXIMO NÓ NEXT
	// A VANTAGEM DA LISTA ENCADEADA SOBRE O ARRAY É PODEMOS ADICIONAR E REMOVER ELEMENTOS SEM REALOCAR MEMORIA

	// criando um lista encadeada simples
	list := LinkedList{}
	list.Append(69)
	list.Append(79)
	list.Append(89)
	list.Append(99)

	fmt.Println("Lista Encadeada Simples: ")
	list.Print()

	list.RemoveFirst()
	fmt.Println("Após remover o primeiro: ")
	list.Print()

	// oque fazemos com os codigos da lista encadeada
	//1. Criamos uma estrutura Node para armazenar os dados.
	//2. Criamos a LinkedList, que gerencia os nós.
	//3. Implementamos funções para:
	//	a. Adicionar elementos (Append)
	//	b. Remover o primeiro elemento (RemoveFirst)
	//	c. Imprimir os elementos (Print)
	//Essa implementação usa uma lista simplesmente encadeada, que não pode
	//voltar para o nó anterior.

}

// Exibe os elementos da lista
func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

// Adiciona um elemento no final da lista
func (l *LinkedList) Append(data int) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Remove o primeiro elemento da lista
func (l *LinkedList) RemoveFirst() {
	if l.head != nil {
		l.head = l.head.next
	}
}

// ESTRUTURA DA LISTA ENCADEADA
type LinkedList struct {
	head *Node
}

// ESTRUTURA DE NÓ EM GO
type Node struct {
	data int
	next *Node
}

// PASSGEM POR VALOR
func increment(val int) {
	val++
	fmt.Println("Dentro da func increment: ", val)
}

// PASSAGEM POR REFERÊNCIA
func incrementobyponteiro(val *int) {
	*val++
	fmt.Println("Dentro da func incrementobyponteiro: ", *val)
}
