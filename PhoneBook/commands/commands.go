package commands

import "fmt"

func AddContact(contact map[string]string, name, namber string) {
	contact[name] = namber
	fmt.Println("Добавлен контакт", name, "с номером", namber)
}

func GetContact(contact map[string]string, name string) {
	if _, ok := contact[name]; ok {
		fmt.Println("Контакт ", contact[name])
	} else {
		fmt.Println("Контакт ", name, "не найден")
	}
}

func DeleteContact(contact map[string]string, name string) {
	if _, ok := contact[name]; ok {
		delete(contact, name)
		fmt.Println("Контакт ", name, "удален")
	} else {
		fmt.Println("Контакт ", name, "не найден")
	}
}
func GetAllContact(contact map[string]string) {

	fmt.Println("Телефонная книга:")
	for name, namber := range contact {
		fmt.Println(name, ": ", namber)
	}

}
