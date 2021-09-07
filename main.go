package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Greet struct {
	Name string
	Time string
}
func main() {
	greet := Greet{"anonymous", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("templates/greet.html"))
	http.Handle("/static/", 
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if name := r.FormValue("name"); name != "" {
			greet.Name = name;
		}

		if err := templates.ExecuteTemplate(w, "greet.html", greet); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8090", nil))
}