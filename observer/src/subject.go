package src

type subject interface {
	RegisterObserver(o Observer)
	Deregister(o Observer)
	NotifyAll()
}
