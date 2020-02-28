package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Api returns json like that.
// {
//   "genres": [
//     {
//       "id": 28,
//       "name": "Action"
//     },
//     {
//       "id": 12,
//       "name": "Adventure"
//     },
//     ]
// }
type genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Genres struct for api
type genres struct {
	Genres []genre `json:"genres"`
}

var movieGenres genres
var tvGenres genres

var movieLink = "https://api.themoviedb.org/3/genre/movie/list?api_key=a585bd999f72b48ddc0dfd46e70a7b80&language=en-US"
var tvLink = "https://api.themoviedb.org/3/genre/tv/list?api_key=a585bd999f72b48ddc0dfd46e70a7b80&language=en-US"

// SaveGenres creating list of genres struct
func SaveGenres() {
	saver(movieLink, &movieGenres)
	saver(tvLink, &tvGenres)
}

func saver(link string, genreType *genres) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("couldn't get given link")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("response body couldn't read")
	}
	json.Unmarshal(body, &genreType)
	fmt.Println("recorded")
}

func findGenre(genreName string, genreType genres) int {
	for _, value := range genreType.Genres {
		if strings.ToLower(value.Name) == genreName {
			return value.ID
		}
	}
	return 0
}
