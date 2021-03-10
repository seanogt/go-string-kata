package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

//Add Adds 2 to numbers
func Add(numbersInput string) (int, error) {
	if len(numbersInput) == 0 {
		return 0, nil
	}
	numbers, err := extractNumbers(numbersInput)
	if err != nil {
		return 0, err
	}
	return totalNumbers(numbers)
}

func extractNumbers(numbersInput string) (numbers []int, err error) {
	var numbersBuf strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(numbersInput))
	scanner.Scan()
	l := scanner.Text()
	delimiters := extractDelimiters(l)
	if l[:2] != "//" {
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
	return appendNumbers(numbers, strings.Split(preppedNums, ","))
}

func extractDelimiters(delimiterInput string) (delimiters []string) {
	if delimiterInput[:2] == "//" {
		delimArray := []rune(delimiterInput[2:])
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
	} else {
		delimiters = []string{","}
	}
	return delimiters
}

func appendNumbers(dest []int, source []string) ([]int, error) {
	for _, c := range source {
		i, err := strconv.Atoi(strings.TrimSpace(c))
		if err != nil {
			return dest, err
		}
		dest = append(dest, i)
	}

	return dest, nil
}

func totalNumbers(numbers []int) (total int, err error) {
	var (
		negativeNumBuff strings.Builder
		negativesFound  bool
	)
	for _, i := range numbers {
		if i < 0 {
			negativesFound = true
			negativeNumBuff.WriteString(fmt.Sprint(i))
			negativeNumBuff.WriteString("\t")
		} else if i <= 1000 {
			total += i
		}
	}
	if negativesFound {
		err = fmt.Errorf("Negatives not allowed. Negatives found:%v", negativeNumBuff.String())
	}
	return total, err
}
