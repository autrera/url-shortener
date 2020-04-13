package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

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

func createShortCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	w.Write([]byte("Creating the short code"))
}

func main() {
	http.HandleFunc("/", handleRootPath)
	http.HandleFunc("/shortcodes", createShortCode)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
