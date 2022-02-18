package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Class string
}

// Options 模式
type Options func(p *Person)

func WithClass(class string) Options {
	return func(p *Person) {
		p.Class = class
	}
}

func NewPerson(name string, age int, options ...Options) *Person {
	p := &Person{
		Name: name,
		Age:  age,
	}
	for _, option := range options {
		option(p)
	}
	return p
}

func main() {
	p1 := NewPerson("张三", 18)
	p2 := NewPerson("李四", 18, WithClass("三班"))

	fmt.Printf("%v\n", p1)
	fmt.Printf("%v\n", p2)
}
