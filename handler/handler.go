package handler

import (
	"golangweb/entity"
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

	data := []entity.Product{
		{ID: 1, Name: "Susu", Price: 12000, Stock: 1200},
		{ID: 2, Name: "Teh", Price: 5000, Stock: 2},
		{ID: 3, Name: "Kopi", Price: 10000, Stock: 1300},
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

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("ini adalah get"))
	case "POST":
		w.Write([]byte("ini adalah post"))
	default:
		http.Error(w, "This is just Bug, just calm :)", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "This is just Bug, just calm :)", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "This is just Bug, just calm :)", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "This is just Bug, just calm :)", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "This is just Bug, just calm :)", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{
			"name":    name,
			"message": message,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "This is just Bug, just calm :)", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "This is just Bug, just calm :)", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "This is just Bug, just calm :)", http.StatusInternalServerError)
}
