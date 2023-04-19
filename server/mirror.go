package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/admin/index", Index)
		http.ListenAndServe(":8080", mux)
	}()

	go func() {
		mux1 := http.NewServeMux()
		mux1.HandleFunc("/admin/index", ServerIndex)
		mux1.HandleFunc("/liaison/callback", ServerTest)
		http.ListenAndServe(":8082", mux1)
	}()
	select {}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("admin/index,hello8080")
	w.Write([]byte("admin/index,hello8080"))
}

func ServerIndex(w http.ResponseWriter, r *http.Request) {
	var result = make(map[string]string)
	r.ParseForm()
	for k, v := range r.PostForm {
		if len(v) < 1 {
			continue
		}
		result[k] = v[0]
	}
	fmt.Println(result)
	fmt.Println("admin/index,hello8082")
	w.Write([]byte("server/index,hello8082"))
}

func ServerTest(w http.ResponseWriter, r *http.Request) {
	var result = make(map[string]string)
	r.ParseForm()
	for k, v := range r.PostForm {
		if len(v) < 1 {
			continue
		}
		result[k] = v[0]
	}
	fmt.Println(result)
	fmt.Println("/test,test8082")
	w.Write([]byte("test,test8082"))
}

func formatRequest(r *http.Request) string {
	var request []string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}
	return strings.Join(request, "\n")
}
