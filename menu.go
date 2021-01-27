package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func MainMenu() {
	ClearConsole()
	reader := bufio.NewReader(os.Stdin)

	for {
		ClearConsole()
		fmt.Println("Поиск персонажей")
		fmt.Println("Выбор папки:       1")
		fmt.Println("Список персонажей: 2")
		fmt.Println("Поиск:             3")
		fmt.Println("Выход:             0")
		fmt.Println("---------------------")
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\r\n", "", -1)

		if text == "0" {
			break
		}

		switch text {
		case "1":
			fmt.Println("Введите путь к папке:")
			fmt.Print("> ")
			charPath, _ = reader.ReadString('\n')
			charPath = strings.Replace(charPath, "\r\n", "", -1)
			CharLoad()
			PressEnter()
		case "2":
			if charList != nil && len(charList) > 0 {
				ShowChars()
				PressEnter()
			} else {
				fmt.Println("Персонажи не загружены")
				PressEnter()
			}
		case "3":
			if charList != nil && len(charList) > 0 {
				Search()
			} else {
				fmt.Println("Персонажи не загружены")
				PressEnter()
			}
		}
	}
}

func CharLoad() {
	files, err := ioutil.ReadDir(charPath)
	if err != nil {
		fmt.Println("Неверный путь", charPath)
		return
	}
	charList = make([]Character, 0)

	for _, f := range files {
		if f.IsDir() {
			_, err := os.Stat(FullDefPath(f.Name()))
			if err != nil {
				if os.IsNotExist(err) {
					//No def file in a folder
					continue
				} else {
					log.Fatal(err)
				}
			}
			//Start reading def file
			ReadDef(f.Name())
		}
	}
	if len(charList) > 0 {
		fmt.Println(len(charList), "персонажей загружено")
	} else {
		fmt.Println("Ничего не найдено!")
	}
}

func ReadDef(s string) {
	file, err := os.Open(FullDefPath(s))
	if err != nil {
		log.Fatal(err)
	}
	bufferedReader := bufio.NewReader(file)
	var name, dname, author string

	for {
		if name != "" && dname != "" && author != "" {
			break
		}
		dataString, err := bufferedReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
		dataString = strings.TrimSpace(dataString)
		if len(dataString) < 6 {
			continue
		}

		if (strings.ToLower(dataString[:6]) == "author") || (strings.ToLower(dataString[:4]) == "name") || (len(dataString) >= 11 && strings.ToLower(dataString[:11]) == "displayname") {
			index := strings.Index(dataString, "\"")
			if index == -1 {
				continue
			}
			index2 := strings.Index(dataString[index+1:len(dataString)], "\"")
			if index2 == -1 {
				continue
			}
			if strings.ToLower(dataString[:6]) == "author" {
				author = dataString[index+1 : index+1+index2]
				continue
			}
			if strings.ToLower(dataString[:4]) == "name" {
				name = dataString[index+1 : index+1+index2]
				continue
			}
			dname = dataString[index+1 : index+1+index2]

		}
	}
	if name != "" {
		charList = append(charList, Character{FileName: s, Name: name, Dname: dname, Author: author})
	}
}

func ShowChars() {
	fmt.Println(len(charList), "chars in the list")
	for _, char := range charList {
		fmt.Printf("Папка: %s; Имя: %s; Отображаемое имя: %s; Автор: %s\n", char.FileName, char.Name, char.Dname, char.Author)
	}
}

func FullDefPath(s string) string {
	return (charPath + "\\" + s + "\\" + s + ".def")
}

func ClearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PressEnter() {
	fmt.Println("Нажмите Enter чтобы продолжить")
	fmt.Scanln()
}
