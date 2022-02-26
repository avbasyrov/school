package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ticker := time.NewTicker(10 * time.Second)
	file, err := os.OpenFile("Vremya",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Unknown file", err)
		os.Exit(1)
	}
	defer file.Close()

	for _ = range ticker.C {
		data := []byte(time.Now().Format("02-01-2006 15:04:05"))
		file.Write(data)
	}
}
