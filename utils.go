package main

import (
	"fmt"
	"os"
)

func writeToLogFile(text string) {
	f, err := os.OpenFile("/git/test.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(text)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	f.Close()
}
