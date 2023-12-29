package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func checkFile(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	file, err := os.Open("test.csv")
	checkFile(err)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	out, err := os.Create("out.ldif")
	checkFile(err)

	outStr := ""
	var restOfLines []string
	fileScanner.Scan()
	firstLine := fileScanner.Text()

	for fileScanner.Scan() {
		restOfLines = append(restOfLines, fileScanner.Text())
	}

	defer file.Close()

	firstLineValues := strings.Split(firstLine, ";")

	for _, line := range restOfLines {
		splitValues := strings.Split(line, ";")

		for i := 0; i < len(firstLineValues); i++ {
			if splitValues[i] != "" {
				outStr += firstLineValues[i] + ": " + splitValues[i] + "\n"
			}
		}
	}

	out.WriteString(outStr)

	defer out.Close()
}
