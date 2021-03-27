package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"

	. "scripts"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "tpl.gohtml", req.Form)

	// Remove clone part below and move it to new script file
	// invode function inside the ssh-into-server.go file
	// url := req.Form["fname"][0]
	Ssh("url")

	// if req.Form["fname"] != nil {
	// 	fmt.Println(req.Form["fname"][0])
	// 	clone(req.Form["fname"][0])
	// }
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
