package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/hujimori/go-chat/data"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, World, %s!", request.URL.Path[1:])
}

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"/templates/layout.html", "templates/navbar.html", "templates/index.html"}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}

func err(w http.ResponseWriter, r *http.Request) {
}

func login(w http.ResponseWriter, r *http.Request) {

}

func logout(w http.ResponseWriter, r *http.Request) {

}

func signup(w http.ResponseWriter, r *http.Request) {

}

func signupAccount(w http.ResponseWriter, r *http.Request) {

}

func authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, _ := data.
}

func newThread(w http.ResponseWriter, r *http.Request) {

}

func createThread(w http.ResponseWriter, r *http.Request) {

}

func postThread(w http.ResponseWriter, r *http.Request) {

}

func readThread(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)

}
