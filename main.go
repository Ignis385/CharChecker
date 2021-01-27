package main

type Character struct {
	FileName string
	Name     string
	Dname    string
	Author   string
}

var charPath string

var charList []Character

func main() {
	MainMenu()
}
