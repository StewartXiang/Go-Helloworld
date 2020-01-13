package main

import (
	"crypto/sha256"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request){
	_ = r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form{
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	_, _ = fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request){
	fmt.Println("method", r.Method)
	if r.Method == "GET"{
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		_ = r.ParseForm()
		if len(r.Form["username"][0]) == 0{
			fmt.Println("No Name!")
		}
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func register(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		t, _ := template.ParseFiles("register.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		_ = r.ParseForm()
		if len(r.Form["username"][0]) == 0{
			fmt.Println("No Name!")
		} else if r.Form["password"][0] != r.Form["password_again"][0]{
			fmt.Println("No Confirm")
		} else {
			h := sha256.New()
			fmt.Println(r.Form["password"][0])
			_, _ = io.WriteString(h, r.Form["password"][0])
			fmt.Printf("%x", h.Sum(nil))
		}
	}
}

func main(){
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	err := http.ListenAndServe(":9090", nil)
	if err != nil{

		log.Fatal("ListenAndServe:", err)

	}
}