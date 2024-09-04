package main

import (
	"log"
	"net/http"

	"html/template"

	"github.com/gorilla/mux"
)

const ROOT = "/Users/mitya/Desktop/stemoly-website/"

func main() {
	log.Fatal(run())
}

func run() error {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(ROOT + "views/index.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, nil)
	})

	r.HandleFunc("/proposal", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(ROOT + "views/proposal.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, nil)
	})

	r.HandleFunc("/team", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(ROOT + "views/team.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, nil)
	})

	r.HandleFunc("/donate", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(ROOT + "views/donate.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, nil)
	})

	r.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(ROOT + "views/contact.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, nil)
	})

	r.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

	return http.ListenAndServe(":8080", r)
}
