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
	var out string
	for j := 0; j < len(totalVals); j++ {
		for i := 0; i < len(keys); i++ {
			out += keys[i] + ": " + totalVals[j][i] + "\n"
		}
	}

	fmt.Println(out)

	csvFile.Close()

}
