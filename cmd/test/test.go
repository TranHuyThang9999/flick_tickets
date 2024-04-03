package main

import "fmt"

type Action interface {
	Run() Action
	Swim() Action
	Fly() Action
}

type Animal struct {
	name string
}

func (a *Animal) Run() Action {
	fmt.Printf("%s is running.\n", a.name)
	return a
}

func (a *Animal) Swim() Action {
	fmt.Printf("%s is swimming.\n", a.name)
	return a
}

func (a *Animal) Fly() Action {
	fmt.Printf("%s is flying.\n", a.name)
	return a
}

type Human struct {
	name string
}

func (h *Human) Run() Action {
	fmt.Printf("%s is running.\n", h.name)
	return h
}

func (h *Human) Swim() Action {
	fmt.Printf("%s is swimming.\n", h.name)
	return h
}

func (h *Human) Fly() Action {
	fmt.Printf("%s is flying.\n", h.name)
	return h
}

func main() {
	animal := &Animal{name: "Bird"}
	human := &Human{name: "John"}

	// Gọi các phương thức liên tiếp
	animal.Run().Swim().Fly()
	human.Run().Swim().Fly()
}
