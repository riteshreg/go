package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("creating")

	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")

	err := f.Close()

	if err != nil {
		panic(err)
	}

}

func main() {
	f := createFile(`C:\Users\Ritesh Khadka/Desktop/x.txt`)
	defer closeFile(f)
	writeFile(f)
}
