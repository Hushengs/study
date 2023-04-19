package main

import (
	"fmt"
	"net/http"
)

func main() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/admin/index", Index)
		http.ListenAndServe(":8080", mux)
	}()

	go func() {
		mux1 := http.NewServeMux()
		mux1.HandleFunc("/server/index", ServerIndex)
		http.ListenAndServe(":8081", mux1)
	}()
	select {}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("admin/index,hello")
	w.Write([]byte("admin/index,hello"))
}

func ServerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server/index,hello")
	w.Write([]byte("server/index,hello"))
}
