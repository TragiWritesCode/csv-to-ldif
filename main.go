package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	csvFile, err := os.Open("test.csv")
	check(err)
	scanner := bufio.NewScanner(csvFile)
	scanner.Scan()

	//https://www.bytesizego.com/blog/reading-file-line-by-line-golang
	keys := strings.Split(scanner.Text(), ";")

	var totalVals [][]string
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), ";")
		totalVals = append(totalVals, vals)
	}
	fileInfo, err := os.Stat("test.csv")

	if err != nil {
		check(err)
	}

	var out strings.Builder
	out.Grow(int(fileInfo.Size()))
	for j := 0; j < len(totalVals); j++ {
		for i := 0; i < len(keys); i++ {
			out.WriteString(strings.Trim(keys[i], " "))
			out.WriteString(": ")
			out.WriteString(strings.Trim(totalVals[j][i], " "))
			out.WriteString("\n")
		}
		out.WriteString("\n")
	}

	fmt.Println(out.String())

	csvFile.Close()

}
