package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

//func IsValidLine(text string) bool {
//***Non-regex solution***
// textLen := len(text)

// //confirm text is at least long enough to start with a tag.
// //check also required before using "text[0:4]" to declare testString
// if textLen < 5 {
// 	return false
// }

// testString := text[0:4]
// tags := []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}
// for _, tag := range tags {
// 	if tag != testString {
// 		return false
// 	}
// }
// return true
//}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-=~*]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	sum := 0

	for _, line := range lines {
		if re.MatchString(line) {
			sum++
		}
	}
	return sum
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	tagged := []string{}

	for _, line := range lines {
		founds := re.FindStringSubmatch(line)
		if founds != nil {
			line = fmt.Sprintf("[USR] %s %s", founds[1], line)
		}
		tagged = append(tagged, line)
	}
	return tagged
}
