package main

import (
	// "fmt"
	// "net/http"
	"os"
	"log"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, "soumya")
	if err != nil {
		log.Fatalln(err)
	}
}