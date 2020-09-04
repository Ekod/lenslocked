package main

import (
	"os"
	"text/template"
)

//User for the template
type User struct {
	Name   string
	Dog    string
	Age    byte
	DogAge byte
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name:   "John Smith",
		Dog:    "Pupps",
		Age:    32,
		DogAge: 5,
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
