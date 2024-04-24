package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type Item struct {
	Name string
	Date time.Time
	Tags string
	Link string
}

var items []Item

func main() {
	defer func() {
		// Завершаем работу с клавиатурой при выходе из функции
		_ = keyboard.Close()
	}()

	fmt.Println("Программа для добавления url в список")
	fmt.Println("Для выхода и приложения нажмите Esc")

OuterLoop:
	for {
		// Подключаем отслеживание нажатия клавиш
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'a':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			// Добавление нового url в список хранения
			fmt.Println("Введите новую запись в формате <url описание теги>")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(text)
			if len(args) < 3 {
				fmt.Println("Введите правильный аргументы в формате url описание теги")
				continue OuterLoop
			}
			// Напишите свой код здесь
			link := args[0]
			name := args[1]
			tags := args[2]
			date := time.Now()

			item := Item{Name: name, Date: date, Tags: tags, Link: link}
			items = append(items, item)

			fmt.Println("URL успешно добавлен в список")

		case 'l':
			// Вывод списка добавленных url. Выведите количество добавленных url и список с данными url
			// Вывод в формате
			// Имя: <Описание>
			// URL: <url>
			// Теги: <Теги>
			// Дата: <дата>
			// Напишите свой код здесь
			if len(items) != 0 {
				fmt.Println("Список добавленных URL:")
				for i, item := range items {
					fmt.Printf("%d. Имя: %s\n   URL: %s\n   Теги: %s\n   Дата: %s\n", i+1, item.Name, item.Link, item.Tags, item.Date.Format(time.DateTime))
				}
			} else {
				fmt.Println("Нет сохранненых URL")
			}
		case 'r':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}
			// Удаление URL из списка хранения
			// Напишите свой код здесь
			fmt.Println("Введите номер URL, который нужно удалить")

			var num int
			fmt.Scan(&num)
			if num < 1 || num > len(items) {
				fmt.Println("Неверный номер URL")
				continue OuterLoop
			}

			items = append(items[:num-1], items[num:]...)
			fmt.Println("URL успешно удален")

			
		default:
			// Если нажата Esc выходим из приложения
			if key == keyboard.KeyEsc {
				return
			}
		}
	}
}
