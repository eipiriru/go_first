package webapps

import (
	"encoding/json"
	"fmt"
	"go_first/kalkulator"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

var rootProj string = "webapps"

func Route(address string) {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path.Join(rootProj, "assets")))))
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/konversiSuhu", konversiSuhu)
	http.HandleFunc("/processFormKonversiSuhu", processFormKonversiSuhu)

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

func konversiSuhu(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join(rootProj, "views", "konversiSuhu.html")
	var tmp, err = template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title": "konversisuhu",
	}
	err = tmp.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func processFormKonversiSuhu(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var suhu, _ = strconv.ParseFloat(r.FormValue("suhu"), 64)
		var dari, _ = strconv.Atoi(r.FormValue("dari"))
		var ke, _ = strconv.Atoi(r.FormValue("ke"))

		var result float64
		result = kalkulator.KonversiSuhu(dari, ke, suhu)

		payload := struct {
			Result float64
		}{
			result,
		}

		jsonInBytes, err := json.Marshal(payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonInBytes)
	default:
		http.Error(w, "", http.StatusBadRequest)
	}
}
