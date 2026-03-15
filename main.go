package main

import (
	"fmt"
	
	"github.com/n1kno/practic2/pkg/list"
	"github.com/n1kno/practic2/pkg/queue"
	"github.com/n1kno/practic2/pkg/stack"
)

func main() {
	// Демонстрация работы стека с разными типами
	demonstrateStack()
	
	// Демонстрация работы очереди с разными типами
	demonstrateQueue()
	
	// Демонстрация работы списка с разными типами
	demonstrateList()
}

func demonstrateStack() {
	fmt.Println("=== Демонстрация стека ===")
	
	// Стек с int
	intStack := stack.New[int]()
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	
	fmt.Println("Стек int:")
	for !intStack.IsEmpty() {
		if val, ok := intStack.Pop(); ok {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println()
	
	// Стек со string
	stringStack := stack.New[string]()
	stringStack.Push("Hello")
	stringStack.Push("World")
	
	if val, ok := stringStack.Peek(); ok {
		fmt.Printf("Верхний элемент: %s\n", val)
	}
}

func demonstrateQueue() {
	fmt.Println("\n=== Демонстрация очереди ===")
	
	// Очередь с int
	intQueue := queue.New[int]()
	intQueue.Enqueue(1)
	intQueue.Enqueue(2)
	intQueue.Enqueue(3)
	
	fmt.Println("Очередь int:")
	for !intQueue.IsEmpty() {
		if val, ok := intQueue.Dequeue(); ok {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println()
}

func demonstrateList() {
	fmt.Println("\n=== Демонстрация списка ===")
	
	// Список с int
	intList := list.New[int]()
	intList.InsertAtEnd(1)
	intList.InsertAtEnd(2)
	intList.InsertAtBeginning(0)
	intList.InsertAtEnd(3)
	
	fmt.Printf("Список int: %s\n", intList.Display(" -> "))
	
	// Список со string
	stringList := list.New[string]()
	stringList.InsertAtEnd("Go")
	stringList.InsertAtEnd("Rust")
	stringList.InsertAtEnd("Python")
	
	fmt.Printf("Список string: %s\n", stringList.Display(" -> "))
}
