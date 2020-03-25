package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Movie struct for json
type Movie struct {
	Title      string   `json:"Title"`
	Year       string   `json:"Year"`
	Runtime    string   `json:"Runtime"`
	Genre      string   `json:"Genre"`
	Director   string   `json:"Director"`
	ImdbRating string   `json:"imdbRating"`
	Ratings    []Rating `json:"Ratings"`
}

// Rating struct for ratings fields in json
type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

var movieObject Movie

// ErrNotFound for if type of search doesn't exist
var ErrNotFound = errors.New("couldn't find")

// SearchMovie search movie using omdb api
func SearchMovie(movieName string, searchType string) error {
	token := os.Getenv("OMDB_APIKEY")
	if token == "" {
		return fmt.Errorf("environment variable couldn't find")
	}
	s := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&t=", token)
	link := s + strings.Replace(movieName, " ", "+", -1)
	movie, err := Unmarshal(link)
	if err != nil {
		return err
	}
	if movie.Title == "" {
		return ErrNotFound
	}
	fmt.Printf("Title: %s\nDirector: %s\nYear: %s\nGenre: %s\n",
		movie.Title, movie.Director, movie.Year, movie.Genre)
	return nil
}

// Unmarshal to given link
func Unmarshal(link string) (Movie, error) {
	resp, err := http.Get(link)
	if err != nil {
		return movieObject, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return movieObject, err
	}
	json.Unmarshal(body, &movieObject)
	return movieObject, nil
}
