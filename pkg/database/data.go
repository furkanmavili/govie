package database

type Service interface {
	Close() error

	CreateTable(listName string) error

	DeleteList(listName string) error
	IsValid(listName string) bool
	ShowListsAll() error
	ShowList(listName string) error
	SaveMovie(movieName, listName string, rate float32) error
	DeleteMovie(movieName, listName string) error
}
