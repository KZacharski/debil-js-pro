package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		filename := os.Args[1]
		fmt.Println(filename)
	} else if len(os.Args) < 2 {
		fmt.Println("Błąd: Nie wystarczająco argumentów. Dodaj nazwę pliku po komendzie, np. 'debil-js-pro debil.djs'.")
	} else {
		fmt.Println("Błąd: Za dużo argumentów. debil-js Pro przyjmuje jedynie jeden argument, np. 'debil-js-pro debil.djs'.")
	}
}
