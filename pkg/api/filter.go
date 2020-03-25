package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type movie struct {
	ID          int     `json:"id"`
	Lang        string  `json:"original_language"`
	Title       string  `json:"title"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
	Date        string  `json:"release_date"`
}

type movieDiscover struct {
	Results []movie `json:"results"`
}

var md movieDiscover

//FilterGenre func getting 20 movies that given genre
func FilterGenre(genreName string) {

	genre := findGenre(genreName, movieGenres)
	if genre == 0 {
		fmt.Println("Aradığınız genre bulunamadı.")
	}
	genreID := strconv.Itoa(genre)
	token := os.Getenv("MOVIEDB_APIKEY")
	if token == "" {
		fmt.Println("moviedb api key env variable couldn't find")
	}
	s := fmt.Sprintf("https://api.themoviedb.org/3/discover/movie?api_key=%s&language=en-US&sort_by=vote_average.desc&include_adult=false&include_video=false&page=1&vote_count.gte=1000&with_genres=", token)
	var link = s + genreID
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("couldn't get given link")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("response body couldn't read")
	}
	json.Unmarshal(body, &md)

	for i, value := range md.Results {
		fmt.Printf("%d.movie: %s, Vote Average: %.2f, ReleaseDate:%s\n", i+1, value.Title, value.VoteAverage, value.Date)
	}
}
