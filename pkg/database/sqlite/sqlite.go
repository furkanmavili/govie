package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/furkanmavili/govie/pkg/database"
	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct {
	db *sql.DB
	tx *sql.Tx
}

func New() (database.Service, error) {
	var db *sql.DB
	var err error
	path := "./govie.db"
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS users(
            userID INTEGER PRIMARY KEY,
            'username' TEXT NOT NULL
        );
         CREATE TABLE IF NOT EXISTS lists(
           listID INTEGER PRIMARY KEY,
           'name' TEXT NOT NULL,
           date TEXT,
           owner NOT NULL REFERENCES users(userID)
         );
        CREATE TABLE IF NOT EXISTS movies(
			movieID INTEGER PRIMARY KEY,
			'name' TEXT NOT NULL,
			rate INTEGER,
			date TEXT,
            listID NOT NULL REFERENCES lists(listID)
		  );`)

	if err != nil {
		_ = db.Close()
		return nil, err
	}
	// Verify connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &sqlite{db, nil}, nil
}

func (s *sqlite) Close() error {
	return s.db.Close()
}

func (s *sqlite) CreateList(listName string) error {
	currentTime := time.Now()
	query := "INSERT INTO lists(name, date, owner) VALUES(?, ?, ?)"
	name := listName
	date := currentTime.Format("01-02-2006")
	owner := 1
	_, err := s.db.Exec(query, name, date, owner)
	return err
}

func (s *sqlite) DeleteList(listName string) error {

	stmt, err := s.db.Prepare("delete from lists where name = ?")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(listName)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlite) IsValid(listName string) bool {
	rows, err := s.db.Query("select name from lists")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		if name == listName {
			return false
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func (s *sqlite) ShowListsAll() error {
	rows, err := s.db.Query("select name, date from lists")
	if err != nil {
		return err
	}
	defer rows.Close()
	fmt.Println("  --------------------- ")
	fmt.Println(" | Lists || Created at |")
	fmt.Println(" |---------------------|")
	for rows.Next() {
		var name string
		var date string

		err = rows.Scan(&name, &date)
		if err != nil {
			return err
		}
		fmt.Printf(">| %s || %s |\n", name, date)
	}
	fmt.Println("  -------------------- ")
	return nil
}

func (s *sqlite) SaveMovie(movieName, listName string, rate float32) error {
	currentTime := time.Now()
	query := "INSERT INTO movies(name, rate, date, listID) VALUES(?, ?, ?, ?)"
	date := currentTime.Format("01-02-2006")
	listID, err := findListID(listName, s)
	if err != nil {
		return err
	}
	if listID == 0 {
		return fmt.Errorf("liste bulunamadı")
	}
	_, err = s.db.Exec(query, movieName, rate, date, listID)
	return err
}

func findListID(listName string, s *sqlite) (int, error) {
	rows, err := s.db.Query("select listID, name from lists")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		var listID int
		var name string
		err = rows.Scan(&listID, &name)
		if err != nil {
			return 0, err
		}
		if listName == name {
			return listID, nil
		}
	}
	return 0, nil

}
