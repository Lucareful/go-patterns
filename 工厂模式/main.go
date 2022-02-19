package main

type Person struct {
	name string
	age  int
}

// NewPerson 新建一个 Person （简单的工厂模式：确保新的实例结构是用所需的参数构造的）
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// PersonInterface 接口,接口允许您在不暴露内部实现的情况下定义行为
type PersonInterface interface {
	GetName() string
}

type person struct {
	name string
	age  int
}

func (p *person) GetName() string {
	return p.name
}

func NewPersonInterface(name string, age int) PersonInterface {
	return &person{
		name: name,
		age:  age,
	}
}

// 工厂功能: 创建单个结构的专用实例时，工厂函数很有用
// Person1
type Person1 struct {
	name string
	age  int
}

// NewPersonFactory 该age属性现在在NewPersonFactory闭包中被抽象出来
func NewPersonFactory(age int) func(name string) Person1 {
	return func(name string) Person1 {
		return Person1{
			name: name,
			age:  age,
		}
	}
}
