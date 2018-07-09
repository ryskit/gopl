package main

import (
	"html/template"
	"os"
	"log"
)

func main() {
	const templ1 = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ1))
	var data struct {
		A string // 信頼されないプレーンテキスト
		B template.HTML // 信頼されたHTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
