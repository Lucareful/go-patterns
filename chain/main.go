package main

type Department interface {
	execute(*Patient)
	setNext(Department) Department
}

func main() {

	reception := &Reception{}

	reception.setNext(&Doctor{}).setNext(&Medical{}).setNext(&Cashier{})

	patient := &Patient{name: "abc"}

	reception.execute(patient)
}
