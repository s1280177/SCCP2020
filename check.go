package main

import (
	"bytes"
	"fmt"
	"net/http"
)

var hashmap map[int]string

func main() {
	http.HandleFunc("/", Hoge)
	http.HandleFunc("/todo", Todo)
	hashmap = make(map[int]string)
	http.ListenAndServe(":8080", nil)
}

func Hoge(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Access/n")
	fmt.Fprint(w, "hoge")
}

func Todo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "show todo\n")
		fmt.Fprint(w, hashmap)
	case http.MethodPost:
		bufbody := new(bytes.Buffer)
		bufbody.ReadFrom(r.Body)
		fmt.Fprint(w, "add todo\n")
	case http.MethodPut:
		fmt.Fprint(w, "updata todo\n")
	case http.MethodDelete:
		fmt.Fprint(w, "delete todo\n")
	default:
		fmt.Fprint(w, "error\n")
	}

}
