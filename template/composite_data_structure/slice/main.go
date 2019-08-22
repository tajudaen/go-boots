package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
var tpl2 *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	tpl2 = template.Must(template.ParseFiles("tpl-02.gohtml"))
}
func main() {
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}

	err02 := tpl2.Execute(os.Stdout, sages)
	if err02 != nil {
		log.Fatalln(err)
	}

}
