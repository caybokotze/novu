package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// - Create HTTP Server
//  |- should respond to server requests on localhost
//  |- should be able to connect to local running mysql instances
//  |- should act as a proxy for sql commands
//  |- should have a built in JWT auth system to handle secure requests from the front end

// - Server Database
//  |- there should be a dedicated server database strictly for user registration / storing tokens and validating users.

// - Create Client project
//  |- the client project should have some UI that caters for the ability to write sql code.
//  |- the ui should also have the ability to execute sql code and display the results in table format.
//  |- the ui should display the results in horizontal datatable format

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
