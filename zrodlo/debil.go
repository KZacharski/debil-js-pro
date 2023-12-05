package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func Debil(filename string) {
	fmt.Println(filename)
	src, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	src = bytes.Replace(src, []byte("konsola"), []byte("console"), -1)
	src = bytes.Replace(src, []byte("dziennik"), []byte("log"), -1)
	newName := strings.Replace(filename, ".djs", ".js", -1)
	if err = os.WriteFile(newName, src, 0666); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) == 2 {
		filename := os.Args[1]
		Debil(filename)
	} else if len(os.Args) < 2 {
		fmt.Println("Błąd: Nie wystarczająco argumentów. Dodaj nazwę pliku po komendzie, np. 'debil-js-pro debil.djs'.")
	} else {
		fmt.Println("Błąd: Za dużo argumentów. debil-js Pro przyjmuje jedynie jeden argument, np. 'debil-js-pro debil.djs'.")
	}
}
