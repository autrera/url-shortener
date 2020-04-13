package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ShortCode struct {
	Id int
	LongUrl string
	ShortUrl string
}

type NewShortCodePayload struct {
	Url string
}

var HumbleStorage []ShortCode

func handleRootPath(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.URL.Path == "/" {
		// Load the index.html page
		body, err := ioutil.ReadFile("index.html")
		if err != nil {
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}
		w.Write([]byte(body))
	} else {
		// Let's search the short url
		shortcode := r.URL.Path
		shortcode = strings.TrimPrefix(shortcode, "/")
		w.Write([]byte(shortcode))
	}
}

func handleNewShortCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Get the body data
	var payload NewShortCodePayload;
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Check if the url is not already registered!
	var registeredUrl ShortCode;
	for _, v := range HumbleStorage {
		if v.LongUrl == payload.Url {
			registeredUrl = v
		}
	}

	// Create new short code
	HumbleStorage = append(HumbleStorage, ShortCode{len(HumbleStorage) + 1, payload.Url, "http://localhost:8080/abcd"})
	fmt.Println(HumbleStorage)
}

func main() {
	http.HandleFunc("/", handleRootPath)
	http.HandleFunc("/shortcodes", handleNewShortCode)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
