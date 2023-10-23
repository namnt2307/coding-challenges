package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	fileName := "" //default is empty to check stdin for data
	action := ""
	actionList := []string{"-c", "-m", "-w"}

	if len(os.Args) == 3 {
		action = os.Args[1]
		fileName = os.Args[2]

	} else if len(os.Args) == 2 {
		if isAction(os.Args[1], actionList) == true {
			action = os.Args[1]
		} else {
			fileName = os.Args[1]
		}
	}
	bs, err := read(fileName)
	if err != nil {
		fmt.Println(err)
	}

	switch action {
	case "-c":
		fmt.Println(countByte(bs))
	case "-l":
		fmt.Println(countLine(bs))
	case "-w":
		fmt.Println(countWord2(bs))
	case "-m":
		fmt.Println(countChar(bs))

	default:
		fmt.Println(countByte(bs), countLine(bs), countWord2(bs))
	}

}

func isAction(a string, actionList []string) bool {
	for _, v := range actionList {
		if a == v {
			return true
		}

	}
	return false
}

func read(fileName string) ([]byte, error) {
	switch fileName {
	case "":
		// fi, err := os.Stdin.Stat()
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// if (fi.Mode() & os.ModeCharDevice) != 0 {
		// 	fmt.Println("Nothing to count")
		// 	os.Exit(1)
		// }
		// fmt.Println(fi.Mode(), os.ModeCharDevice)

		return io.ReadAll(os.Stdin)
	default:
		return os.ReadFile(fileName)
	}
}

func countByte(bs []byte) int {
	return len(bs)
}

func countLine(bs []byte) int {
	lineCount := 0
	for _, val := range bs {
		if val == 10 {
			lineCount += 1
		}
	}
	return lineCount
}

func countWord(bs []byte) int {
	textString := string(bs)
	regex, err := regexp.Compile("([^ \n\r\t]+)") // \r is for CR (13)
	if err != nil {
		fmt.Println(err)
	}

	wordCounts := len(regex.FindAllString(textString, -1))

	return wordCounts
}

// countWord or countWord2 use regex to remove space, line feed, and tab, CR.
func countWord2(bs []byte) int {
	// c := 0
	s := string(bs)
	regex := regexp.MustCompile("(\\S+)") // \r is for CR (13)

	return len(regex.FindAllString(s, -1))
}

func countChar(bs []byte) int {
	s := string(bs)
	return len([]rune(s))
}

// condition
// length == 1
//     run_all from stdin

// length == 2
//     if flag[1] == cmd {
//         run from stdin
//     }
//     if flag[1] == file {
//         run_all from file
//     }

// length == 3
//     -> cmd := flag[1]
//     -> file := flag[2]
//     -> run
