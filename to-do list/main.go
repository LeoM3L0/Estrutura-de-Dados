package main

import (
	"fmt"
)

func main() {
	fmt.Println("H3LL000")

	lista := TaskList{}

	// adicionando tarefas
	lista.AddTask("Lavar roupa")
	lista.AddTask("Capinar lote")
	lista.AddTask("ouvir racionais")

	lista.ShowTasks()
	lista.CompleteTask()
	lista.ShowTasks()
	lista.CompleteTask()
	lista.ShowTasks()
}

// Estrtura do nó
type Task struct {
	name string // Nome da tarefa
	next *Task  // ponteiro para proxima tarefa
}

// Estrutura lista encadeada
type TaskList struct {
	head *Task // ponteiro para o inicio da lista
}

// metodo para adicionar uma nova tarefa a lista
func (t *TaskList) AddTask(name string) {
	newTask := &Task{name: name}

	if t.head == nil {
		t.head = newTask // se a lsita estiver vazia adiciona como primeira
	} else {
		current := t.head
		for current.next != nil {
			current = current.next
		}
		current.next = newTask // adiciona ao final
	}
}

// metodo para remover a primeira tarefa da lista
func (t *TaskList) CompleteTask() {
	if t.head == nil {
		fmt.Println("Lista de Tarefa Vazia")
		return
	}

	fmt.Printf(" Tarefa Concluida: %s\n", t.head.name)
	t.head = t.head.next // remove a primeira tarefa
}

// metodo para exibir todas as tarefas pendentes
func (t *TaskList) ShowTasks() {
	if t.head == nil {
		fmt.Println(" Nenhuma Tarefa pendente")
		return
	}

	fmt.Println("Tarefas Pendentes: ")
	current := t.head
	for current != nil {
		fmt.Printf("- %s\n", current.name)
		current = current.next
	}
}
