package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func (fl *FirstLinear) ReadFile() {
	// Open the file.
	f, _ := os.Open(fl.Filename)
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	fl.Testnum, _ = strconv.Atoi(scanner.Text())
	fl.Engine = make(map[int][]string, fl.Testnum)
	fl.Query = make(map[int][]string, fl.Testnum)
	for i := 0; i < fl.Testnum; i++ {
		scanner.Scan()
		enginenum, _ := strconv.Atoi(scanner.Text())
		fl.Enginenum = append(fl.Enginenum, enginenum)

		var bufengine []string
		for x := 0; x < enginenum; x++ {
			scanner.Scan()
			bufengine = append(bufengine, strings.Trim(scanner.Text(), " "))
		}
		fl.Engine[i] = bufengine

		scanner.Scan()
		querynum, _ := strconv.Atoi(scanner.Text())
		fl.Querynum = append(fl.Querynum, querynum)

		var bufquery []string
		for x := 0; x < querynum; x++ {
			scanner.Scan()
			bufquery = append(bufquery, strings.Trim(scanner.Text(), " "))
		}
		fl.Query[i] = bufquery
	}
}
