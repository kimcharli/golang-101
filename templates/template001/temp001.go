package main

import (
	"text/template"
	"os"
)

type Inventory struct {
	Material string
	Count int
}

func main() {
	sweaters := Inventory{"Wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil { panic(err) }
	
}


