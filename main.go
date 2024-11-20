package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/image", imageHandler)
	http.HandleFunc("/greet", greetHandler)
	fmt.Println("Server is running & listening on port http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	imgData, err := ioutil.ReadFile("life.jpg")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(imgData)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Ishwari")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	response := fmt.Sprintf("Namaskar %s", name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
