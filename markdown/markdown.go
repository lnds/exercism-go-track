package markdown

// implementation to refactor

import (
	"fmt"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	header := 0
	// step 2: functions for strong and em
	markdown = strong(markdown)
	markdown = emphasis(markdown)
	pos := 0
	list := 0
	html := ""
	for pos < len(markdown) {
		char := markdown[pos]
		// step 1 change if with switch
		switch char {
		case '#':
			for char == '#' {
				header++
				pos++
				char = markdown[pos]
			}
			html += fmt.Sprintf("<h%d>", header)
			pos++
			continue
		case '*':
			if list == 0 {
				html += "<ul>"
			}
			html += "<li>"
			list++
			pos += 2
			continue
		case '\n':
			if list > 0 {
				html += "</li>"
			}
			if header > 0 {
				html += fmt.Sprintf("</h%d>", header)
				header = 0
			}
			pos++
			continue
		default:
			html += string(char)
			pos++
		}
	}
	if header > 0 {
		return html + fmt.Sprintf("</h%d>", header)
	}
	if list > 0 {
		return html + "</li></ul>"
	}
	return encloseWithTag(html, "p")

}

func strong(markdown string) string {
	result := strings.Replace(markdown, "__", "<strong>", 1)
	return strings.Replace(result, "__", "</strong>", 1)
}

func emphasis(markdown string) string {
	result := strings.Replace(markdown, "_", "<em>", 1)
	return strings.Replace(result, "_", "</em>", 1)
}

func encloseWithTag(text, tag string) string {
	return fmt.Sprintf("<%s>%s</%s>", tag, text, tag)
}
