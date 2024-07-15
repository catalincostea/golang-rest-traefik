package main

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Movie struct {
	Title string `json:"title"`
}

type ApiResponse struct {
	Results []Movie `json:"results"`
}

// main 
func main() {
	fmt.Println("Starting main..")
	// Create Gin router
	router := gin.Default()

	// Register Routes
	router.GET("/movies", getAllMovies)

	// Start the server
	router.Run(":3000")
}

// TODO: Modify this function
func getAllMovies(c *gin.Context) {
	// c.String(http.StatusOK, "TODO")

	url := "https://swapi.dev/api/films/"

	resp, err := http.Get(url)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make request"})
			return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Unexpected status code"})
		return
	}

	var apiResponse ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode JSON"})
			return
	}

	var titles []string
	for _, movie := range apiResponse.Results {
			titles = append(titles, movie.Title)
	}

	c.JSON(http.StatusOK, gin.H{"titles": titles})
}


func test_getAllMovies(c *gin.Context) {
	// testing with by local file

	fmt.Println("Starting getAllMovies..")
	data, err := ioutil.ReadFile("response_file.json")
	if err != nil {
		fmt.Println("file error")
	}

	c.Data(http.StatusOK, "application/json", data)
}