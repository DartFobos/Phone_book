package main

import (
	"bufio"
	"fmt"
	"os"
)

func menu() {
	fmt.Println("Желаете добавить контакт: введите 1")
	fmt.Println("Желаете найти контакт: введите 2")
	fmt.Println("Желаете удалить контакт: введите 3")
	fmt.Println("Желаете вывести все контакты: введите 4")
	fmt.Println("Выйти: введите 5")
}

func addContact(contact map[string]string, name, namber string) {
	contact[name] = namber
	fmt.Println("Добавлен контакт", name, "с номером", namber)
}

func getContact(contact map[string]string, name string) {
	if _, ok := contact[name]; ok {
		fmt.Println("Контакт ", contact[name])
	} else {
		fmt.Println("Контакт ", name, "не найден")
	}
}

func deleteContact(contact map[string]string, name string) {
	if _, ok := contact[name]; ok {
		delete(contact, name)
		fmt.Println("Контакт ", name, "удален")
	} else {
		fmt.Println("Контакт ", name, "не найден")
	}
}
func getAllContact(contact map[string]string) {

	fmt.Println("Контакты:")
	for name, namber := range contact {
		fmt.Println(name, ": ", namber)
	}

}

func main() {
	contact := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		menu()
		fmt.Print("Введите команду: ")

		scanner.Scan()
		text := scanner.Text()
		switch text {
		case "5":
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неверный команда")
		case "1":
			var name string
			var namber string
			fmt.Print("Введите имя: ")

			scanner.Scan()
			name = scanner.Text()
			fmt.Print("Введите номер: ")
			scanner.Scan()
			namber = scanner.Text()
			addContact(contact, name, namber)
		case "2":

			fmt.Print("Введите имя: ")

			scanner.Scan()
			name := scanner.Text()
			getContact(contact, name)
		case "3":

			fmt.Print("Введите имя: ")

			scanner.Scan()
			name := scanner.Text()
			deleteContact(contact, name)
		case "4":
			getAllContact(contact)
		}
	}
}
