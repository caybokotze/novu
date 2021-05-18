package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", homePage)
	_ = http.ListenAndServe(":8080", nil)
}
type PageVariables struct {
	Date string
	Time string
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	var HomePageVars = PageVariables{
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}
	t, err := template.ParseFiles("home.html")
	if err != nil {
		log.Print("template parsing error, ", err)
	}
	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Print("template executing error, ", err)
	}
}
