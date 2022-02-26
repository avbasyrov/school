package main

import (
	"fmt"
	"os"
	"time"
)

/*Написать программу, которая раз в 10 секунд добавляет в текстовый
файл новую строчку с текущим временем.
Одновременно с этим, программка запрашивает на экране от пользователя ввод строки.
Как только пользователь
что-то вводит и нажимает enter - добавлять введенную пользователем строку в тот же файлик,
куда время пишется.
И потом снова давать пользователю что-нибудь ввести. И так по кругу
*/
func main() {
	file, err := os.OpenFile("Vremya",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Unknown file", err)
		os.Exit(1)
	}
	defer file.Close()

	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for _ = range ticker.C {
			data := []byte(time.Now().Format("02-01-2006 15:04:05") + "\n")
			file.Write(data)
		}
	}()

	var name string
	for {
		fmt.Print("Введите имя:")
		fmt.Fscan(os.Stdin, &name)
		fmt.Println(name)
		file.Write([]byte(name + "\n"))
	}
}
