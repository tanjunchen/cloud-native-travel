package main

import "fmt"

type Animal interface {
	Move()

	Shout()
}

type Dog struct {
}

func (dog Dog) Move() {
	fmt.Println("A dog moves with its logs.")
}

func (dog Dog) Shout() {
	fmt.Println("wang wang wang.")
}

type Bird struct {
}

func (bird Bird) Move() {
	fmt.Println("A bird moves with its logs.")
}

func (bird Bird) Shout() {
	fmt.Println("bird shouts")
}

func main() {
	var animal Animal

	animal = Dog{}

	animal.Move()
	animal.Shout()

	animal = Bird{}
	animal.Move()
	animal.Shout()

}
