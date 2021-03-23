package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "tpl.gohtml", req.Form)
	if req.Form["fname"] != nil {
		fmt.Println(req.Form["fname"][0])
		clone(req.Form["fname"][0])
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./src/views/tpl.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8000", d)
}

func clone(repo string) {
	cmd := exec.Command("git", "clone", repo)
	err := cmd.Run()
	if err != nil {
		fmt.Println("something went wrong")
	}
}
