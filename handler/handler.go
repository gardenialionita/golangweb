package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//w.Write([]byte("Homepage"))

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"content": "Learning Golang Web with XXX",
	}

	tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario World"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	//fmt.Fprintf(w, "Product page: %d", idNumb)
	data := map[string]interface{}{
		"content": idNumb,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}
}
