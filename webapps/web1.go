package webapps

import (
	"fmt"
	"net/http"
)

func Route(address string) {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("webapps/assets"))))

	fmt.Println(http.Dir("/webapps/assets"))
	fmt.Println("Server started on", address)

	server := new(http.Server)
	server.Addr = address

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	msg := "Hello ini Index"
	w.Write([]byte(msg))
}

func hello(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	w.Write([]byte(msg))
}
