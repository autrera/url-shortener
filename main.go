package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"errors"
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

func getLongUrlFromShortUrl(shortUrl string) (string, error) {
	for _, v := range HumbleStorage {
		if shortUrl == v.ShortUrl {
			return v.LongUrl, nil
		}
	}
	return "", errors.New("Long url not found")
}

func getCompressedUrlFromLongUrl(longUrl string) (CompressedUrl, error) {
	for _, v := range HumbleStorage {
		if v.LongUrl == longUrl {
			return v, nil
		}
	}
	return CompressedUrl{0, "", ""}, errors.New("Compressed url not found")
}

func compressUrl(longUrl string) (CompressedUrl, error) {
	newId := len(HumbleStorage) + 1
	newShortUrlCode := generateShortUrl(newId)

	compressedUrl := CompressedUrl{newId, longUrl, newShortUrlCode}
	return compressedUrl, nil
}

func storeCompressedUrl(compressedUrl CompressedUrl) () {
	HumbleStorage = append(HumbleStorage, compressedUrl)
	fmt.Println(HumbleStorage)
}

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
		LongUrl, err := getLongUrlFromShortUrl(shortUrl)
		if err != nil {
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, LongUrl, 301)
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
	compressedUrl, err := getCompressedUrlFromLongUrl(payload.Url)

	if compressedUrl.ShortUrl == "" {
		// Generate the short code for this url
		compressedUrl, _ = compressUrl(payload.Url)
		storeCompressedUrl(compressedUrl)
	}

	// Send the short url for the url received
	js, err := json.Marshal(NewShortUrlPayload{ Url: compressedUrl.ShortUrl })
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", handleRootPath)
	http.HandleFunc("/short-urls", handleNewShortUrl)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
