package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
	}

}

func main() {
	http.HandleFunc("/", SayHello)
	http.HandleFunc("/login", Login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("listenAndServe:", err)
	}
}
