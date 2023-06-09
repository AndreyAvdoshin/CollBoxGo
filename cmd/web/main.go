package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	//Логирование
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//Преобразование путей к обработчикам
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/collbox", collbox)
	mux.HandleFunc("/collbox/create", create)

	// Получаем доступ к статическим файлам
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/web/static")})
	mux.Handle("/static", http.NotFoundHandler())
	// Отрезаем "/static" из пути
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//Инициализиуер структуру для логирования ошибок HTTP
	srv := &http.Server{
		Addr:     ":8000",
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// Запустили сервер и проверили на ошибки
	infoLog.Printf("Сервер запущен")
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	// Если открыть файл не удалось возвращяем ошибку
	if err != nil {
		return nil, err
	}
	// Проверяем если в пути директория и подменяем на 404
	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}
	return f, nil
}
