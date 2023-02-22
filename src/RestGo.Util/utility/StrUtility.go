package utility

import (
	"strings"
	"unicode/utf8"
)

type RestString string

func SplitCommand(inputStr string) (cmd string, para string) {
	inSingleQuote := false
	for i, c := range inputStr {
		if c == '\'' {
			inSingleQuote = !inSingleQuote
		} else if !inSingleQuote && c == ' ' {
			cmd = strings.Trim(inputStr[:i], "'")  // Trim single quotes from cmd if present
			para = strings.TrimSpace(inputStr[i:]) // Trim whitespace from para
			break
		}
	}

	if cmd == "" { // If cmd is empty, set it to the original input string
		cmd = inputStr
	}
	return // Return the command and parameter strings
}

func (s *RestString) Replace(old string, new string) *RestString {
	str := string(*s)
	str = strings.ReplaceAll(str, old, new)
	*s = RestString(str)
	return s
}

func (s *RestString) String() string {
	return string(*s)
}

// 每多少字換行
func InsertNewLine(s string, args ...int) string {
	maxLen := 25
	if len(args) > 0 {
		maxLen = args[0]
	}
	var visibleLen int
	var result string
	s = strings.ReplaceAll(s, "\r\n", "\n")
	var inEscapeSeq bool
	for i := 0; i < len(s); i++ {
		if s[i] == '\u001B' {
			inEscapeSeq = true
			result += string(s[i])
			continue
		}
		if inEscapeSeq {
			if s[i] == 'm' {
				inEscapeSeq = false
			}
			result += string(s[i])
			continue
		}
		if s[i] == '\n' {
			visibleLen = 0
			result += "\n"
			continue
		}
		r, size := utf8.DecodeRuneInString(s[i:])
		if utf8.RuneLen(r) == 2 {
			visibleLen += 2
		} else {
			visibleLen++
		}
		if visibleLen > maxLen {
			visibleLen = 0
			result += "\n"
		}
		result += string(r)
		i += size - 1
	}
	return result
}
