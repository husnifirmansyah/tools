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

func (sl *SecondLinear) ReadFile() {
	f, _ := os.Open(sl.Filename)
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	sl.Testnum, _ = strconv.Atoi(scanner.Text())
	sl.FirstQuery = make(map[int][]string, sl.Testnum)
	sl.SecondQuery = make(map[int][]string, sl.Testnum)
	for i := 0; i < sl.Testnum; i++ {
		scanner.Scan()
		allQueryNum, _ := strconv.Atoi(scanner.Text())
		sl.AllQueryNum = append(sl.AllQueryNum, allQueryNum)

		scanner.Scan()
		arrQuery := strings.Split(scanner.Text(), " ")
		firstQueryNum, _ := strconv.Atoi(arrQuery[0])
		secondQueryNum, _ := strconv.Atoi(arrQuery[1])
		sl.FirstQueryNum = append(sl.FirstQueryNum, firstQueryNum)
		sl.SecondQueryNum = append(sl.SecondQueryNum, secondQueryNum)

		var bufFirstQuery []string
		for x := 0; x < firstQueryNum; x++ {
			scanner.Scan()
			bufFirstQuery = append(bufFirstQuery, scanner.Text())
		}
		sl.FirstQuery[i] = bufFirstQuery
		var bufSecondQuery []string
		for x := 0; x < secondQueryNum; x++ {
			scanner.Scan()
			bufSecondQuery = append(bufSecondQuery, scanner.Text())
		}
		sl.SecondQuery[i] = bufSecondQuery
	}
}
