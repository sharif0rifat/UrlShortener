package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	println("Server is Starting")
	// http.ListenAndServe(":9090", nil)
	r := mux.NewRouter()
	r.HandleFunc("/", HandleGet).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}
func HandleGet(w http.ResponseWriter, r *http.Request) {
	var name string = r.URL.Query().Get("url")
	w.Write([]byte(name + "   " + RandStringRunes()))
}

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes() string {
	b := make([]rune, 7)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
