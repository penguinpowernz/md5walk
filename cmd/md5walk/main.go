package main

import (
	"log"
	"os"

	"github.com/penguinpowernz/md5walk"
)

func main() {
	sums, err := md5walk.Walk(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	_, err = sums.Write(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
