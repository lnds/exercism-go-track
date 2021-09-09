package markdown

// implementation to refactor

import (
	"fmt"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	// step 2: functions for strong and em
	markdown = strong(markdown)
	markdown = emphasis(markdown)

	lines := strings.Split(markdown, "\n")
	listPos := -1
	for i, line := range lines {
		switch {
		// step 4 process headers
		case strings.HasPrefix(line, "#"):
			lines[i] = parseHeader(line)
		case strings.HasPrefix(line, "*"):
			lines[i] = parseListItem(line)
			if listPos < 0 {
				listPos = i
			}
		default:
			lines, listPos = fixList(lines, listPos, i), -1
			lines[i] = encloseWithTag(line, "p")
		}
	}
	lines = fixList(lines, listPos, len(lines)-1)
	return strings.Join(lines, "")

}

// final step fixList
func fixList(lines []string, listPos, i int) []string {
	if listPos >= 0 {
		lines[listPos] = "<ul>" + lines[listPos]
		lines[i] = lines[i] + "</ul>"
	}
	return lines
}

// step 2 add function for emphasis and strong
func strong(markdown string) string {
	result := strings.Replace(markdown, "__", "<strong>", 1)
	return strings.Replace(result, "__", "</strong>", 1)
}

func emphasis(markdown string) string {
	result := strings.Replace(markdown, "_", "<em>", 1)
	return strings.Replace(result, "_", "</em>", 1)
}

// step 3 add encloseWithTag
func encloseWithTag(text, tag string) string {
	return fmt.Sprintf("<%s>%s</%s>", tag, text, tag)
}

func parseHeader(line string) string {
	header := 0
	for strings.HasPrefix(line, "#") {
		header++
		line = line[1:]
	}
	line = line[1:]
	htag := fmt.Sprintf("h%d", header)
	return encloseWithTag(line, htag)
}

func parseListItem(line string) string {
	for strings.HasPrefix(line, "*") {
		line = line[1:]
	}
	line = line[1:]
	return encloseWithTag(line, "li")
}
