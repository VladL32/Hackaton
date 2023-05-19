package trash

import (
	"html/template"
	"net/http"
)

var (
	repoo PostRepository = NewPostRepository()
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/home.html")
	Posts, _ := repoo.FindAll()
	tmpl.Execute(w, Posts)
}

func Person(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/person.html")
	Posts, _ := repoo.FindAll()
	post := Posts[0]

	tmpl.Execute(w, post)
}
