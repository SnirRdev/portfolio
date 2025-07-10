package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type UserInput struct {
	Name   string   `json:"name"`
	Bio    string   `json:"bio"`
	Skills []string `json:"skills"`
}

const dataFile = "data.json"

var (
	formTmpl      = template.Must(template.ParseFiles("templates/form.html"))
	portfolioTmpl = template.Must(template.ParseFiles("templates/portfolio.html"))
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", showForm)
	http.HandleFunc("/generate", handleForm)
	http.HandleFunc("/all", showAll)
	http.HandleFunc("/json", showJSON)
	http.HandleFunc("/edit", editForm)

	log.Println("📡 Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func showForm(w http.ResponseWriter, r *http.Request) {
	// Send default empty skills so that index works in template
	formTmpl.Execute(w, UserInput{
		Skills: []string{"", "", ""},
	})
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	user := UserInput{
		Name:   r.FormValue("name"),
		Bio:    r.FormValue("bio"),
		Skills: []string{
			r.FormValue("skill1"),
			r.FormValue("skill2"),
			r.FormValue("skill3"),
		},
	}

	var users []UserInput
	if _, err := os.Stat(dataFile); err == nil {
		bytes, _ := ioutil.ReadFile(dataFile)
		json.Unmarshal(bytes, &users)
	}

	// Replace existing user if name matches
	var updated []UserInput
	for _, u := range users {
		if u.Name != user.Name {
			updated = append(updated, u)
		}
	}
	updated = append(updated, user)

	data, _ := json.MarshalIndent(updated, "", "  ")
	ioutil.WriteFile(dataFile, data, 0644)

	portfolioTmpl.Execute(w, user)
}

func showAll(w http.ResponseWriter, r *http.Request) {
	var users []UserInput
	if _, err := os.Stat(dataFile); err == nil {
		bytes, _ := ioutil.ReadFile(dataFile)
		json.Unmarshal(bytes, &users)
	}
	portfolioTmpl.Execute(w, users)
}

func showJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := os.ReadFile(dataFile)
	if err != nil {
		http.Error(w, "Failed to read JSON", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func editForm(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	var users []UserInput
	var selected UserInput

	if _, err := os.Stat(dataFile); err == nil {
		bytes, _ := ioutil.ReadFile(dataFile)
		json.Unmarshal(bytes, &users)
		for _, u := range users {
			if u.Name == name {
				selected = u
				break
			}
		}
	}

	// Ensure skills has 3 elements
	for len(selected.Skills) < 3 {
		selected.Skills = append(selected.Skills, "")
	}

	formTmpl.Execute(w, selected)
}
