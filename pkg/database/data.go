package database

type Service interface {
	Close() error
	CreateList(listName string) error
	DeleteList(listName string) error
	IsValid(listName string) bool
	ShowListsAll() error
	SaveMovie(movieName, listName string, rate float32) error
}
