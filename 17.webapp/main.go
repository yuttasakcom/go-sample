package main

import (
	"net/http"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// Solution1
	// f, err := os.Open("public" + r.URL.Path)

	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	log.Println(err)
	// }

	// defer f.Close()

	// var contentType string
	// switch {
	// case strings.HasSuffix(r.URL.Path, "css"):
	// 	contentType = "text/css"
	// case strings.HasSuffix(r.URL.Path, "html"):
	// 	contentType = "text/html"
	// case strings.HasSuffix(r.URL.Path, "png"):
	// 	contentType = "text/png"
	// default:
	// 	contentType = "text/plain"
	// }

	// w.Header().Add("Content-Type", contentType)
	// io.Copy(w, f)

	// Solution 2
	// 	http.ServeFile(w, r, "public"+r.URL.Path)
	// })

	http.ListenAndServe(":8000", http.FileServer(http.Dir("public")))
}
