package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

/*
Написать программу, которая раз в 10 секунд добавляет в текстовый
файл новую строчку с текущим временем.
Одновременно с этим, программка запрашивает на экране от пользователя ввод строки.
Как только пользователь
что-то вводит и нажимает enter - добавлять введенную пользователем строку в тот же файлик,
куда время пишется.
И потом снова давать пользователю что-нибудь ввести. И так по кругу

*/

func main() {

	c := make(chan os.Signal)
	signal.Notify(c)
	ticker := time.NewTicker(10 * time.Second)
	stop := make(chan bool)

	go func() {
		defer func() { stop <- true }()
		for {
			select {
			case <-ticker.C:

				now := time.Now()
				fmt.Println(now.Format("02-01-2006 15:04:05")) //шаблонный

				var name string
				fmt.Println("Как тебя зовут?")
				fmt.Scanf("%s\n", &name)

				now = time.Now()
				fmt.Println(now.Format("02-01-2006 15:04:05")) //шаблонный
				var age int
				fmt.Println("Сколько тебе лет?")
				fmt.Scanf("%d\n", &age)

				now = time.Now()
				fmt.Println(now.Format("02-01-2006 15:04:05")) //шаблонный
				var sex string
				fmt.Println("Пойдем в кино?")
				fmt.Scanf("%s\n", &sex)

			case <-stop:
				return
			}
		}
	}()
	// default, чтобы посылать и получать данные без блокировок
	// Блокировка, пока не будет получен сигнал
	<-c
	ticker.Stop()

	// Остановка горутины
	//stop <- true
}
