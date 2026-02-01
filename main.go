package main

import (
	"crud/connection/data/data"
	crud "crud/connection/data/mysqldb"
	"html/template"
	"net/http"
)

const baseDir = "templates/"

func index(w http.ResponseWriter, req *http.Request) {
	var AllColuns = data.AllColunsSenhas{
		All: crud.SelectAllPass(),
	}
	tmpl := template.Must(template.ParseFiles(baseDir + "index.html"))
	tmpl.Execute(w, AllColuns)
}
func newPassw(w http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		req.ParseForm()
		senha_local := req.Form.Get("senha_local")
		senha := req.Form.Get("senha")
		crud.InsertPassword(senha_local, senha)

		http.Redirect(w, req, "/", 0)
	}
	tmpl := template.Must(template.ParseFiles(baseDir + "newPassword.html"))
	tmpl.Execute(w, nil)
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/string/", http.FileServer(http.Dir("/static/"))))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/new", newPassw)
	http.ListenAndServe(":8090", mux)
}
