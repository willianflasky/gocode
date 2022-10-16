package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/register", registerHandle)
	http.HandleFunc("/login", loginHandle)

	err = http.ListenAndServe("127.0.0.1:8888", nil)
	if err != nil {
		fmt.Println("booting error: ", err)
	}
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(500)
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		ok, err := queryUser(username, password)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		if !ok {
			w.Write([]byte("user or password error!"))
			return
		} else {
			w.Write([]byte("welcome login!"))
		}
	} else {
		t, err := template.ParseFiles("./login.html")
		if err != nil {
			w.WriteHeader(500)
		}
		t.Execute(w, nil)
	}
}

func registerHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(500)
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		err = createUser(username, password)
		if err != nil {
			w.WriteHeader(500)
		}
		http.Redirect(w, r, "/register", 302)
	} else {
		t, err := template.ParseFiles("./register.html")
		if err != nil {
			w.WriteHeader(500)
		}
		t.Execute(w, nil)
	}
}
