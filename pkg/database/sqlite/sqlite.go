package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/furkanmavili/govie/pkg/database"
	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct {
	db *sql.DB
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
		`CREATE TABLE IF NOT EXISTS lists(
           listID INTEGER PRIMARY KEY,
           'name' TEXT NOT NULL,
           date TEXT
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
	return &sqlite{db}, nil
}

func (s *sqlite) Close() error {
	return s.db.Close()
}

// CreateTable method is creat table and table name is listName. this method is DONE.
func (s *sqlite) CreateTable(listName string) error {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		movieID INTEGER PRIMARY KEY,
		'movieName' TEXT NOT NULL,
		rate INTEGER,
		date TEXT,
		listID NOT NULL REFERENCES lists(listID)
	  );
	  INSERT INTO lists(name, date) VALUES(?, ?);`, listName)
	currentTime := time.Now().Format("01-02-2006")
	_, err := s.db.Exec(query, listName, currentTime)
	return err
}

func (s *sqlite) DeleteList(listName string) error {

	stmt, err := s.db.Prepare("delete from lists where name = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(listName)
	if err != nil {
		return err
	}
	return nil
}

// it shows all lists with name and date
// TODO: add how much movie consist also print properly.
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

type movie struct {
	name string
	rate int
	date string
}

// ShowList prints context of given list. this method is DONE.
func (s *sqlite) ShowList(listName string) error {
	query := fmt.Sprintf("SELECT movieName, rate, date from %s", listName)
	rows, err := s.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	var movies []movie
	longest := 0
	for rows.Next() {
		var name string
		var rate int
		var date string

		err = rows.Scan(&name, &rate, &date)
		if err != nil {
			return err
		}
		if len(name) > longest {
			longest = len(name)
		}
		m := movie{name, rate, date}
		movies = append(movies, m)
	}

	si := strings.Repeat("-", longest)
	fmt.Println("-------------------" + si)
	for i, v := range movies {
		l := strings.Repeat(" ", longest-len(v.name))
		fmt.Printf("> %d. %s "+l+"%d★   %s\n", i+1, v.name, v.rate, v.date)
	}
	fmt.Println("-------------------" + si)
	return nil
}

// SaveMovie saves the movie to given list. this method is DONE.
func (s *sqlite) SaveMovie(movieName, listName string, rate float32) error {
	err := checkMovie(s, movieName, listName)
	if err != nil {
		return err
	}
	currentTime := time.Now()
	query := fmt.Sprintf("INSERT INTO %s(movieName, rate, date, listID) VALUES(?, ?, ?, ?)", listName)
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

// TODO: verilen listedeki movieyi sil.interface'e eklemeyi unutma.
func (s *sqlite) DeleteMovie(movieName, listName string) error {
	return nil
}

// checkMovie checks the movieName in given list, return error if its exist.
func checkMovie(s *sqlite, movieName, listName string) error {
	query := fmt.Sprintf("SELECT movieName from %s", listName)
	rows, err := s.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return err
		}
		if name == movieName {
			return fmt.Errorf("%s already exist in %s", movieName, listName)
		}
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

// IsValid checks the given listName in lists table
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

// findListID finds id of given list
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
