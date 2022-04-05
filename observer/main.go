package main

import src "github.com/luenci/go-patterns/observer/src"

func main() {

	shirtItem := src.NewItem("Nike Shirt")

	observerFirst := &src.Customer{Id: "abc@gmail.com"}
	observerSecond := &src.Customer{Id: "xyz@gmail.com"}

	shirtItem.RegisterObserver(observerFirst)
	shirtItem.RegisterObserver(observerSecond)

	shirtItem.NotifyAll()
}
