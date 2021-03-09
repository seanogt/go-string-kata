package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the calculator")
}

//Add Adds 2 to numbers
func Add(numbersInput string) (total int, err error) {

	var (
		delimiters      []string
		numbers         []int
		numbersBuf      strings.Builder
		negativeNumBuff strings.Builder
		negativesFound  bool
	)

	negativeNumBuff.WriteString("negatives not allowed. Negatives found: ")
	if len(numbersInput) == 0 {
		return 0, nil
	}

	scanner := bufio.NewScanner(strings.NewReader(numbersInput))

	scanner.Scan()
	l := scanner.Text()
	if numbersInput[:2] == "//" {
		delimiters = extractDelimiters(l[2:])
	} else {
		delimiters = []string{","}
		numbersBuf.WriteString(l)
	}

	for scanner.Scan() {
		l := scanner.Text()
		numbersBuf.WriteString(l)
	}

	preppedNums := numbersBuf.String()
	for _, d := range delimiters {
		preppedNums = strings.ReplaceAll(preppedNums, d, ",")
	}
	numbers = addToSliceIfNumber(numbers, strings.Split(preppedNums, ","))

	for _, i := range numbers {

		if i < 0 {
			negativesFound = true
			negativeNumBuff.WriteString(fmt.Sprint(i))
			negativeNumBuff.WriteString("\t")
		} else if i <= 1000 {
			total += int(i)
		}
	}
	if negativesFound {
		err = fmt.Errorf(negativeNumBuff.String())
	}
	return total, err
}

func extractDelimiters(delimiterInput string) (delimiters []string) {
	delimArray := []rune(delimiterInput)
	currentDelimiter := []rune{}
	for _, c := range delimArray {

		switch c {
		case '[':
			continue
		case ']':
			delimiters = append(delimiters, string(currentDelimiter))
			currentDelimiter = []rune{}
		default:
			currentDelimiter = append(currentDelimiter, c)
		}
	}
	if len(currentDelimiter) != 0 {
		delimiters = append(delimiters, string(currentDelimiter))
	}

	return delimiters
}

func addToSliceIfNumber(dest []int, source []string) []int {
	for _, c := range source {
		i, err := strconv.Atoi(strings.TrimSpace(c))
		if err == nil {
			dest = append(dest, i)
		}
	}

	return dest
}
