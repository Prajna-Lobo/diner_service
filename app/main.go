package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFromFile(filename string, linebreak int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	log, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	lines := strings.Split(string(log), "\n")
	lineMap := make(map[string]bool)
	for i, line := range lines {
		if exists, _ := lineMap[line]; exists {
			return fmt.Errorf("duplicate entry")
		}

		fmt.Println(line)
		if i > linebreak {
			break
		}

		lineMap[line] = true
	}

	return nil
}

func main() {
	filePath := flag.String("fpath", "log.csv", "file path")
	flag.Parse()

	err := ReadFromFile(*filePath, 2)
	if err != nil {
		fmt.Println("Something went wrong")
		fmt.Println(err.Error())
	}
}
