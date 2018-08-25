package main 

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
)

func main() { 
	/* http.HandleFunc("/", sayhelloName) // setting router rule */
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":8080", nil) // setting listening port
    if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
