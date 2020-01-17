package main

import (
	"io/ioutil"
	"log"
)

var a, b = swap("Hello", "World")

func main() {
	defer log.Print("Defer")
	log.Printf("%s %s. Wow I can add: %d\n", a, b, add(100, 200))
	fileName := "README.md"
	doIExist := checkForFileExists(fileName)
	if doIExist == nil {
		log.Print("File Exists")
	} else {
		log.Fatal(doIExist)
	}
}

func swap(input, input2 string) (string, string) {
	return input2, input
}

func checkForFileExists(filePath string) error {
	_, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return nil
}

func add(x, y int) int {
	return x + y
}