package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Обработчик главной страницы
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/web/html/home.page.html",
		"./ui/web/html/base.layout.html",
		"./ui/web/html/footer.partial.html",
	}

	tempParse, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = tempParse.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	//w.Write([]byte("Home Page!"))
}

// Обработчик страниц с записями
func collbox(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	log.Println(id)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Отображение определенной заметки - ID %d...", id)
}

// Обработчик создания новой записи
func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод запрещен", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("Create new record"))
}
