package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	data := []byte(time.Now().Format("02-01-2006 15:04:05"))
	file, err := os.Create("Vremya")
	if err != nil {
		fmt.Println("Unknown file", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write(data)

	fmt.Println("Time is done")
}
