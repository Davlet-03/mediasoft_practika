package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	Name     string
	Age      int
	Position string
	Salary   int
}

var commands = `
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
`

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func readInt(scanner *bufio.Scanner) int {
	for {
		text := readString(scanner)

		number, err := strconv.Atoi(text)
		if err == nil {
			return number
		}

		fmt.Println("Введите число:")
	}
}

func main() {
	const size = 512

	empls := [size]*Employee{}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(commands)
		fmt.Print("\nВведите команду: ")

		cmd := readInt(scanner)

		switch cmd {
		case 1:
			// Добавляем нового сотрудника
			empl := new(Employee)

			fmt.Print("Имя: ")
			empl.Name = readString(scanner)

			fmt.Print("Возраст: ")
			empl.Age = readInt(scanner)

			fmt.Print("Позиция: ")
			empl.Position = readString(scanner)

			fmt.Print("Зарплата: ")
			empl.Salary = readInt(scanner)

			added := false

			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					added = true
					fmt.Println("Сотрудник добавлен!")
					break
				}
			}

			if !added {
				fmt.Println("Нет свободных мест!")
			}

		case 2:
			// Удаляем сотрудника
			fmt.Print("Введите номер сотрудника: ")

			index := readInt(scanner)

			if index < 0 || index >= size {
				fmt.Println("Неверный номер!")
				continue
			}

			if empls[index] == nil {
				fmt.Println("Сотрудника с таким номером нет!")
				continue
			}

			empls[index] = nil

			fmt.Println("Сотрудник удалён!")

		case 3:
			// Выводим список сотрудников
			fmt.Println("Список сотрудников:")

			for i := 0; i < size; i++ {
				if empls[i] != nil {
					fmt.Printf(
						"[%d] Имя: %s, Возраст: %d, Позиция: %s, Зарплата: %d\n",
						i,
						empls[i].Name,
						empls[i].Age,
						empls[i].Position,
						empls[i].Salary,
					)
				}
			}

		case 4:
			fmt.Println("Выход из программы.")
			return

		default:
			fmt.Println("Такой команды нет!")
		}
	}
}
