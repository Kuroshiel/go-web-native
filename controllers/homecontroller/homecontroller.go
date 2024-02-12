package homecontroller

import (
	"net/http"
	"text/template"
)

// Homepage Golang CRUD Web

// Handler
func Welcome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
