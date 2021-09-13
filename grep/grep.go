package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Search(pattern string, flags, files []string) []string {
	result := []string{}
	lineNumbers := false
	names := false
	ignoreCase := false
	invert := false
	matchEntireLine := false
	for _, flag := range flags {
		switch flag {
		case "-n":
			lineNumbers = true
		case "-l":
			names = true
		case "-i":
			ignoreCase = true
		case "-v":
			invert = true
		case "-x":
			matchEntireLine = true
		}
	}
	multi := len(files) > 1
	for _, file := range files {
		result = append(result, searchInFile(pattern, lineNumbers, names, ignoreCase, invert, matchEntireLine, multi, file)...)
	}
	return result
}

func searchInFile(pattern string, n, l, i, v, x, m bool, fileName string) []string {
	result := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		return result
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++
		if find(line, pattern, i, x) == !v {
			result = append(result, format(fileName, line, lineNumber, n, l, m))
			if l {
				break
			}
		}
	}
	return result
}

func find(text, pattern string, i, x bool) bool {
	line := text
	patt := pattern
	if i {
		line = strings.ToUpper(line)
		patt = strings.ToUpper(patt)
	}
	if x {
		return line == patt
	}
	return strings.Contains(line, patt)
}

func format(file, line string, lineNumber int, n, l, m bool) string {
	result := line
	if l {
		result = file
	} else if n {
		result = fmt.Sprintf("%d:%s", lineNumber, line)
	}
	if m && !l {
		result = fmt.Sprintf("%s:%s", file, result)
	}
	return result
}
