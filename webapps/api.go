package webapps

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func api_init() {
	students_init()
	http.HandleFunc("/student", Student)
	http.HandleFunc("/student/", Student2)
	http.HandleFunc("/student/{id}", Student2)
}

func Login(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()

	if !ok {
		w.Write([]byte("Error"))
		return false
	}

	if !auth(username, password) {
		w.Write([]byte("Username or Password is Wrong"))
		return false
	}

	return true
}

func Student(w http.ResponseWriter, r *http.Request) {
	if !Login(w, r) {
		return
	}

	var tempId = r.URL.Query().Get("id")
	var id, _ = strconv.ParseInt(tempId, 10, 64)
	OutputJson(w, getStudents(id))
}

func Student2(w http.ResponseWriter, r *http.Request) {
	if !Login(w, r) {
		return
	}
	var tempId = r.PathValue("id")
	var id, _ = strconv.ParseInt(tempId, 10, 64)
	OutputJson(w, getStudents(id))
}

func OutputJson(w http.ResponseWriter, o []StudentMaster) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
