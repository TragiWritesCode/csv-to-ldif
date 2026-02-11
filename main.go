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
	fmt.Println(keys)

	var totalVals [][]string
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), ";")
		totalVals = append(totalVals, vals)
		totalVals = append(totalVals, vals)
	}
	fmt.Println(keys)
	fmt.Println(totalVals)
	var out string
	for i := 0; i < len(keys); i++ {
		out += keys[i] + ": " + totalVals[i][i] + "\n"
	}

	fmt.Println(out)

	csvFile.Close()

}
