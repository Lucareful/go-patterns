package src

type Subject interface {
	RegisterObserver(o Observer)
	Deregister(o Observer)
	NotifyAll()
}
