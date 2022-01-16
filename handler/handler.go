package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "This is just Bug, just calm :)", http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"content": "Belajar dasar golangweb",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "This is just Bug, just calm :)", http.StatusInternalServerError)
		return
	}

}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumber, err := strconv.Atoi(id)

	if err != nil || idNumber < 1 {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "This is just Bug, just calm :)", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"content": idNumber,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "This is just Bug, just calm :)", http.StatusInternalServerError)
		return
	}

}
