package webapps

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

var rootProj string = "webapps"

func Route(address string) {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path.Join(rootProj, "assets")))))
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)

	fmt.Println("Server started on", address)

	server := new(http.Server)
	server.Addr = address

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join(rootProj, "views", "index.html")
	var tmp, err = template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Apa 1",
		"name":  "Hello World Building",
	}

	err = tmp.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	w.Write([]byte(msg))
}
