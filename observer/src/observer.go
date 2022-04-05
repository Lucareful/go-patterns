package src

type Observer interface {
	Update(string)
	GetID() string
}
