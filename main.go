package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type CompressedUrl struct {
	Id int
	LongUrl string
	ShortUrl string
}

type NewShortUrlPayload struct {
	Url string
}

var HumbleStorage []CompressedUrl

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
		shortUrl := strings.TrimPrefix(r.URL.Path, "/")
		for _, v := range HumbleStorage {
			if shortUrl == v.ShortUrl {
				http.Redirect(w, r, v.LongUrl, 301)
				return
			}
		}

		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
}

func handleNewShortUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Get the body data
	var payload NewShortUrlPayload;
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Check if the url is not already registered!
	var registeredUrl CompressedUrl;
	for _, v := range HumbleStorage {
		if v.LongUrl == payload.Url {
			registeredUrl = v
		}
	}

	if registeredUrl.ShortUrl == "" {
		// Generate the short code for this url
		newId := len(HumbleStorage) + 1
		newShortUrlCode := generateShortUrl(newId)

		// Store the new short code
		registeredUrl = CompressedUrl{newId, payload.Url, newShortUrlCode}
		HumbleStorage = append(HumbleStorage, registeredUrl)
	}

	// Send the short url for the url received
	js, err := json.Marshal(NewShortUrlPayload{ Url: registeredUrl.ShortUrl })
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	fmt.Println(HumbleStorage)
}

func main() {
	http.HandleFunc("/", handleRootPath)
	http.HandleFunc("/short-urls", handleNewShortUrl)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
