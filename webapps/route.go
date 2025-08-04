package webapps

import (
	"fmt"
	"net/http"
	"path"
)

func Route(address string) {
	api_init()
	web1_init()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path.Join(rootProj, "assets")))))
	fmt.Println("Server started on", address)

	server := new(http.Server)
	server.Addr = address

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
