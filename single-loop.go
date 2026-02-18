package main

import (
	"bufio"
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
	var out string
	i := 0
	for j := 0; j < len(vals); j++ {
		out += keys[i] + ": " + vals[j] + "\n"
		i++
		if i == len(keys) {
			i = 0
		}
	}

	//	fmt.Println(out)

	csvFile.Close()

}
