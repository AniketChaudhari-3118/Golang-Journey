package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}
func main() {

	//tpl, err := template.ParseFiles("tpl.gohtml")
	//tpl, err := template.ParseGlob("*.gohtml")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// nf, err := os.Create("index.html")
	// if err != nil {
	// 	log.Println("error creating file", err)
	// }
	// defer nf.Close()

	// err := tpl.Execute(os.Stdout, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//err = tpl.Execute(nf, nil)

	// err = tpl.Execute(os.Stdout, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// tpl, err = tpl.ParseFiles("two.gohtml", "vespa.gohtml")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err := tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}

	// err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

}
