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
	keys := strings.Split(strings.Trim(scanner.Text(), " "), ";")

	var vals []string
	for scanner.Scan() {
		vals = append(vals, strings.Split(strings.Trim(scanner.Text(), " "), ";")...)
	}

	fileInfo, err := os.Stat("test.csv")

	if err != nil {
		check(err)
	}

	var out strings.Builder
	out.Grow(int(fileInfo.Size()))

	i := 0
	for j := 0; j < len(vals); j++ {
		out.WriteString(keys[i])
		out.WriteString(": ")
		out.WriteString(vals[j])
		out.WriteString("\n")
		i++
		if i == len(keys) {
			i = 0
			out.WriteString("\n")
		}
	}

	fmt.Println(out.String())

	csvFile.Close()

}
