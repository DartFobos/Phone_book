package menu

import (
	"fmt"

	"github.com/fatih/color"
)

func Menu() {
	fmt.Print("Желаете записать контакт: ")
	color.Green("нажмите 1")
	fmt.Print("Желаете найти контакт: ")
	color.Green("нажмите 2")
	fmt.Print("Желаете удалить контакт: ")
	color.Green("нажмите 3")
	fmt.Print("Открыть телефонную книгу: ")
	color.Green("нажмите 4")
	fmt.Print("Заершить работу программы: ")
	color.Green("нажмите 5")
}
