package main

import (
	"Phone_book/commands"
	"Phone_book/menu"
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	contact := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		menu.Menu()
		fmt.Print("Введите команду: ")
		scanner.Scan()
		text := scanner.Text()
		switch text {
		case "5":
			fmt.Println("Выход из программы.")
			return
		case "1":
			var name string
			var namber string
			fmt.Print("Введите имя: ")
			scanner.Scan()
			name = scanner.Text()
			fmt.Print("Введите номер: ")
			scanner.Scan()
			namber = scanner.Text()
			commands.AddContact(contact, name, namber)
		case "2":
			fmt.Print("Введите имя: ")
			scanner.Scan()
			name := scanner.Text()
			commands.GetContact(contact, name)
		case "3":
			fmt.Print("Введите имя: ")
			scanner.Scan()
			name := scanner.Text()
			commands.DeleteContact(contact, name)
		case "4":
			commands.GetAllContact(contact)
		default:
			color.New(color.FgHiRed).Print("Внимание! ")
			fmt.Print("Вы ввели неверную команду!\n")
		}
	}
}
