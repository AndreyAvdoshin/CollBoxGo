package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Обработчик главной страницы
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/web/html/home.page.html",
		"./ui/web/html/base.layout.html",
		"./ui/web/html/footer.partial.html",
	}

	tempParse, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
		return
	}

	err = tempParse.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
	}

	//w.Write([]byte("Home Page!"))
}

// Обработчик страниц с записями
func (app *application) collbox(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	log.Println(id)
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Отображение определенной заметки - ID %d...", id)
}

// Обработчик создания новой записи
func (app *application) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create new record"))
}
