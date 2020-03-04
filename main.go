package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var posts []post

type post struct {
	Time time.Time
	Body string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", postHandler).Methods(http.MethodPost)
	r.HandleFunc("/", getHandler).Methods(http.MethodGet)


	fmt.Println("listening for POSTS on 0.0.0.0:3000")
	if err := http.ListenAndServe("0.0.0.0:3000", r); err != nil {
		panic(err)
	}

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	en := json.NewEncoder(w)
	en.Encode(posts)
	return
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	typeHeader := r.Header.Get("X-Gitlab-Event")

	defer r.Body.Close()

	posts = append(posts, post{
		Time: time.Now(),
		Body: string(body),
	})

	fmt.Println("##################START REQUEST##################")
	fmt.Printf("TIME: %v\n", time.Now().String())
	fmt.Printf("MESSAGE TYPE: %v\n", typeHeader)
	fmt.Println(string(body))
	fmt.Println("##################END REQUEST##################")
}
