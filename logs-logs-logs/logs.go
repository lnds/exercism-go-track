package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Message extracts the message from the provided log line.
func Message(line string) string {
	parts := strings.Split(strings.TrimSpace(line), ":")
	if len(parts) == 2 {
		return strings.TrimSpace(parts[1])
	}
	return line
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	msg := Message(line)
	return utf8.RuneCountInString(msg)
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	parts := strings.Split(strings.TrimSpace(line), ":")
	if len(parts) == 2 {
		level := strings.TrimPrefix(parts[0], "[")
		level = strings.TrimSuffix(level, "]")
		return strings.ToLower(level)
	}
	return line
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	level := LogLevel(line)
	msg := Message(line)
	return fmt.Sprintf("%s (%s)", msg, level)
}
