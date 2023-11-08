package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		read(os.Stdin, false)

	} else if len(os.Args) == 2 {
		switch os.Args[1] {
		case "-":
			read(os.Stdin, false)
		case "-n":
			read(os.Stdin, true)
		case "-b":
			read(os.Stdin, false)
		default:
			read(os.Args[1], false)
		}

	} else if len(os.Args) > 2 {
		switch os.Args[1] {
		case "-n":
			for _, v := range os.Args[2:] {
				read(v, true)
			}
		case "-b":
			for _, v := range os.Args[2:] {
				read(v, false)
			}
		default:
			for _, v := range os.Args[1:] {
				read(v, false)
			}
		}
	}

}

func read(file interface{}, nb bool) {
	var bs []byte
	var err error
	switch t := file.(type) {
	case string:
		bs, err = os.ReadFile(file.(string))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case *os.File:
		bs, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Printf("%T", t)
		fmt.Println("type wtf")
	}
	s := string(bs)

	splittedString := strings.Split(s, "\n")
	n := 1
	for _, v := range splittedString {
		if nb == true {
			fmt.Println(n, v)
			n++
		} else {
			if v == "" {
				fmt.Println(v)
				continue
			}
			fmt.Println(n, v)
			n++
		}
	}
}
