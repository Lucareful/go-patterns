package src

import "fmt"

type Item struct {
	Observers []Observer
	Name      string
	InStock   bool
}

func NewItem(name string) *Item {
	return &Item{
		Name: name,
	}
}

func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.Name)
	i.InStock = true
	i.NotifyAll()
}

func (i *Item) RegisterObserver(o Observer) {
	i.Observers = append(i.Observers, o)
}

func (i *Item) Deregister(o Observer) {
	i.Observers = removeFromslice(i.Observers, o)
}

func (i *Item) NotifyAll() {
	for _, observer := range i.Observers {
		observer.Update(i.Name)
	}
}

func removeFromslice(observers []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observers)
	for i, observer := range observers {
		if observerToRemove.GetID() == observer.GetID() {
			observers[observerListLength-1], observers[i] = observers[i], observers[observerListLength-1]
			return observers[:observerListLength-1]
		}
	}
	return observers
}
