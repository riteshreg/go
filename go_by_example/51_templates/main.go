package main

import (
	"os"
	"text/template"
)

func main() {

	t1 := template.New("t1")

	t1 = template.Must(t1.Parse("Value is: {{.}}\n"))

	t1.Execute(os.Stdout, "hi my name is ritesh")

	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

}
