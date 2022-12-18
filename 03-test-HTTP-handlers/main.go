package main

import (
	"net/http"
)

func handleGetFoo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("FOO"))
}

func main() {
	http.HandleFunc("/foo", handleGetFoo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}