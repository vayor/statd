package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	var file *os.File
	if _, err := os.Stat("/etc/statd/statd.yml"); err == nil {
		file, err = os.Open("/etc/statd/statd.yml")
		if err != nil {
			log.Fatalf("[ERROR]: %s\n", err.Error())
		}
	} else {
		if _, err = os.Stat("./statd.yml"); err == nil {
			file, err = os.Open("./statd.yml")
			if err != nil {
				log.Fatalf("[ERROR]: %s\n", err.Error())
			}
		}
	}

	fmt.Println(file)

}
