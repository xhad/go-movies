package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	apiKey = os.Getenv("API_KEY")

	// ErrorBackend when something goes wrong
	ErrorBackend = errors.New("Something went wrong")
)

// Request strucutre
type Request struct {
	ID int `json:"id"`
}

// MovieDBResponse map structure
type MovieDBResponse struct {
	Movies []Movie `json:"results"`
}

// Movie object stucture
type Movie struct {
	Title       string `json:"title"`
	Description string `json:"overview"`
	Cover       string `json:"poster_path"`
	ReleaseDate string `json:"release_date"`
}

// Handler takes movie_genre paramater, returns MovieDB query
func Handler(request Request) ([]Movie, error) {
	url := fmt.Sprintf("http://api.themoviedb.org/3/discover/movie?api_key=%s", apiKey)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []Movie{}, ErrorBackend
	}

	if request.ID > 0 {
		q := req.URL.Query()
		q.Add("with_genres", strconv.Itoa(request.ID))
		req.URL.RawQuery = q.Encode()
	}

	resp, err := client.Do(req)
	if err != nil {
		return []Movie{}, ErrorBackend
	}

	var data MovieDBResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return []Movie{}, ErrorBackend
	}

	return data.Movies, nil
}

func main() {
	fmt.Print(apiKey)
	lambda.Start(Handler)
}
