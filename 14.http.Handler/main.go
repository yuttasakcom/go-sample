package main

import "net/http"

// type myHandler struct {
// 	greeting string
// }

// func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(fmt.Sprintf("%v world", mh.greeting)))
// }

// func main() {
// 	http.Handle("/", &myHandler{greeting: "Hello"})
// 	http.ListenAndServe(":8000", nil)
// }

func main() {
	http.Handle("/", &indexHandler{})
	http.ListenAndServe(":8000", nil)
}

type indexHandler struct{}

func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
