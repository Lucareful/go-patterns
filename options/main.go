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

// uber 推荐的模式暴露接口

type Person2 struct {
	Name  string
	Age   int
	Class string
}

type Options2 interface {
	apply(*Person2)
}

type Class string

func (c Class) apply(p *Person2) {
	p.Class = string(c)
}

func WithClass2(class string) Options2 {
	return Class(class)
}

func NewPerson2(name string, age int, options ...Options2) *Person2 {
	p := &Person2{
		Name: name,
		Age:  age,
	}
	for _, option := range options {
		option.apply(p)
	}
	return p
}

// 更结构化的 options 模式

type Person3 struct {
	Name  string
	Age   int
	Class string
}

type Option3 func(*Person3)

func defaultOption() Person3 {
	return Person3{
		Name:  "luenci",
		Age:   18,
		Class: "一班",
	}
}

func WithName3(name string) Option3 {
	return func(p *Person3) {
		p.Name = name
	}
}

func WithAge3(age int) Option3 {
	return func(p *Person3) {
		p.Age = age
	}
}

func WithClass3(class string) Option3 {
	return func(p *Person3) {
		p.Class = class
	}
}

func NewPerson3(options ...Option3) *Person3 {
	p := defaultOption()
	for _, option := range options {
		option(&p)
	}
	return &p
}

// --------------------------------------------------------------------------------

func main() {
	p1 := NewPerson("张三", 18)
	p2 := NewPerson("李四", 18, WithClass("三班"))
	p3 := NewPerson2("张三", 18)
	p4 := NewPerson2("李四", 18, WithClass2("三班"))

	p5 := NewPerson3()
	p6 := NewPerson3(WithName3("张三"), WithAge3(24))

	fmt.Printf("%v\n", p1)
	fmt.Printf("%v\n", p2)
	fmt.Printf("%v\n", p3)
	fmt.Printf("%v\n", p4)
	fmt.Printf("%v\n", p5)
	fmt.Printf("%v\n", p6)

	// &{张三 18
	// &{李四 18 三班}
	// &{张三 18 }
	// &{李四 18 三班}
	// &{luenci 18 一班}
	// &{张三 24 一班}
}
