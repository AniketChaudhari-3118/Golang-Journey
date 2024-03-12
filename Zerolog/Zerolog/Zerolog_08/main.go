package main

import (
	"Zerolog_08/logger"
	"bytes"
	"html/template"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"

	// "github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

var tpl template.Template

type Search struct {
	Query      string          // Search query string
	Results    *SearchResponse // Search results
	TotalPages int             // Total number of pages based on search results
	NextPage   int             // Next page number
}

// SearchResult represents a single search result from Wikipedia.
type SearchResult struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	// Add more fields as needed
}

// SearchResponse represents the response from the Wikipedia search.
type SearchResponse struct {
	Query struct {
		SearchInfo struct {
			TotalHits int `json:"totalhits"`
		} `json:"searchinfo"`
		Search []SearchResult `json:"search"`
	} `json:"query"`
}

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		l := logger.Get()
		start := time.Now()

		defer func() {
			l.Info().Str("method", req.Method).
				Str("url", req.URL.RequestURI()).
				Str("user_agent", req.UserAgent()).
				Dur("elapsed_ms", time.Since(start)).
				Msg("incoming request")
		}()
		next.ServeHTTP(res, req)
	})
}

func searchHandler(res http.ResponseWriter, req *http.Request) error {
	u, err := url.Parse(req.URL.String())
	if err != nil {
		return err
	}

	params := u.Query()
	searchQuery := params.Get("q")
	pageNum := params.Get("page")
	if pageNum == "" {
		pageNum = "1"
	}

	l := zerolog.Ctx((req.Context()))

	l.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("search_query", searchQuery).Str("page_num", pageNum)
	})

	l.Info().Msgf("Incoming search Query %s on page %s", searchQuery, pageNum)

	nextPage, err := strconv.Atoi(pageNum)
	if err != nil {
		return err
	}

	pageSize := 20
	resultsOffset := (nextPage - 1) * pageSize

	searchResponse, err := func() (*SearchResponse, error) {
		var _ int = resultsOffset
		return searchWikipedia(pageSize)
	}()
	if err != nil {
		return err
	}

	l.Debug().Interface("wikipedia_search_response", searchResponse).Send()

	totalHits := searchResponse.Query.SearchInfo.TotalHits

	search := &Search{
		Query:      searchQuery,
		Results:    searchResponse,
		TotalPages: int(math.Ceil(float64(totalHits) / float64(pageSize))),
		NextPage:   nextPage + 1,
	}

	buf := &bytes.Buffer{}
	err = tpl.Execute(buf, search)
	if err != nil {
		return err
	}

	_, err = buf.WriteTo(res)
	if err != nil {
		return err
	}

	l.Trace().Msgf("search query '%s' succeeded without errors", searchQuery)

	return nil
}

// searchWikipedia simulates a search against Wikipedia.
func searchWikipedia(pageSize int) (*SearchResponse, error) {
	// Simulate fetching search results from Wikipedia.
	// In a real implementation, this would involve making an HTTP request to the Wikipedia API or using some other method.
	// Here, we'll just return some mock data.

	// Assuming we have some mock data for demonstration.
	// Here, we'll create a dummy response with hardcoded data.
	// You would replace this with actual code to fetch data from Wikipedia.

	// For demonstration purposes, we'll create some dummy search results.
	// In a real-world scenario, you would fetch search results from Wikipedia's API or database.
	searchResults := []SearchResult{
		{Title: "Dummy Result 1", URL: "https://en.wikipedia.org/wiki/Dummy_Result_1"},
		{Title: "Dummy Result 2", URL: "https://en.wikipedia.org/wiki/Dummy_Result_2"},
		// Add more search results as needed
	}

	// Create a mock SearchResponse.
	searchResponse := &SearchResponse{
		Query: struct {
			SearchInfo struct {
				TotalHits int `json:"totalhits"`
			} `json:"searchinfo"`
			Search []SearchResult `json:"search"`
		}{
			SearchInfo: struct {
				TotalHits int `json:"totalhits"`
			}{
				TotalHits: len(searchResults), // Total number of search results
			},
			Search: searchResults, // Actual search results
		},
	}

	// Simulating some delay or computation involved in fetching search results.
	// In a real-world scenario, you might have some network latency or processing time.
	time.Sleep(1 * time.Second)

	return searchResponse, nil
}

func searchHandlerWrapper(w http.ResponseWriter, r *http.Request) {
	err := searchHandler(w, r)
	if err != nil {
		// Handle the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Create a new instance of the mux router
	router := mux.NewRouter()

	//use requestLogger as middleware
	router.Use(requestLogger)

	// Register the searchHandler function to handle requests to the "/search" endpoint
	router.HandleFunc("/search", searchHandlerWrapper).Methods("GET")

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", router)
}
