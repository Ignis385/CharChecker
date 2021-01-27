package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Search() {
	ClearConsole()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Искать по...")
	fmt.Println("Имени:               1")
	fmt.Println("Отображаемому имени: 2")
	fmt.Println("Автору:              3")
	fmt.Println("Назад:               0")
	fmt.Println("---------------------")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\r\n", "", -1)

		if text == "0" {
			break
		}

		switch text {
		case "1":
			{
				fmt.Println("Введите имя")
				name, _ := reader.ReadString('\n')
				name = strings.Replace(name, "\r\n", "", -1)
				FilterByName(strings.ToLower(name))
				PressEnter()
				return
			}
		case "2":
			{
				fmt.Println("Please input display name")
				dname, _ := reader.ReadString('\n')
				dname = strings.Replace(dname, "\r\n", "", -1)
				FilterByDName(strings.ToLower(dname))
				PressEnter()
				return
			}
		case "3":
			{
				fmt.Println("Please input author name")
				aname, _ := reader.ReadString('\n')
				aname = strings.Replace(aname, "\r\n", "", -1)
				FilterByAuthor(strings.ToLower(aname))
				PressEnter()
				return
			}
		}
	}
}

func FilterByName(name string) {
	newList := make([]Character, 0)
	for _, ch := range charList {
		if strings.Contains(strings.ToLower(ch.Name), name) {
			newList = append(newList, ch)
		}
	}
	if len(newList) > 0 {
		charList = newList
		fmt.Println(len(charList), "персонажей найдено")
	} else {
		fmt.Println("Ничего не найдено!")
	}

}

func FilterByDName(dname string) {
	newList := make([]Character, 0)
	for _, ch := range charList {
		if strings.Contains(strings.ToLower(ch.Dname), dname) {
			newList = append(newList, ch)
		}
	}
	if len(newList) > 0 {
		charList = newList
		fmt.Println(len(charList), "персонажей найдено")
	} else {
		fmt.Println("Ничего не найдено!")
	}
}

func FilterByAuthor(aname string) {
	newList := make([]Character, 0)
	for _, ch := range charList {
		if strings.Contains(strings.ToLower(ch.Author), aname) {
			newList = append(newList, ch)
		}
	}
	if len(newList) > 0 {
		charList = newList
		fmt.Println(len(charList), "персонажей найдено")
	} else {
		fmt.Println("Ничего не найдено!")
	}
}
