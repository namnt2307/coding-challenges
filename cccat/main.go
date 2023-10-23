package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	switch os.Args[1] {
	case "-n":
		readByLine(os.Args[2:])
	case "-b":
		readByLineLF(os.Args[2:])
	default:
		read(os.Args[1:])
	}

}

func read(fList []string) {

	for _, file := range fList {
		switch file {
		case "-":
			io.Copy(os.Stdout, os.Stdin)
		default:
			f, err := os.Open(file)
			if err != nil {
				fmt.Println(err)
			}
			io.Copy(os.Stdout, f)
		}
	}
}
func readByLineLF(fList []string) {

}

func readByLine(fList []string) {

	for _, file := range fList {
		switch file {
		case "-":
			io.Copy(os.Stdout, os.Stdin)
		case "":
			io.Copy(os.Stdout, os.Stdin)
		default:
			bs, err := os.ReadFile(file)
			if err != nil {
				fmt.Println(err)
			}
			printTextByLine(bs)

		}
	}
}

func printTextByLine(bs []byte) {
	s := string(bs)
	trim := strings.TrimSpace(s)
	splittedString := strings.Split(trim, "\n")
	for n := 0; n < len(splittedString); n++ {
		fmt.Println(n+1, splittedString[n])
	}
}
