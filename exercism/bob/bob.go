/*
Bob is a lackadaisical teenager. In conversation, his responses are very limited.

Bob answers 'Sure.' if you ask him a question, such as "How are you?".

He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).

He answers 'Calm down, I know what I'm doing!' if you yell a question at him.

He says 'Fine. Be that way!' if you address him without actually saying anything.

He answers 'Whatever.' to anything else.

Bob's conversational partner is a purist when it comes to written communication and always follows normal rules regarding sentence punctuation in English.
*/
package bob

import (
	"strings"
	"unicode"
)

//Hey responds with a lackadaisic teenagager response from Bob
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	//checking if remark contains letters as only letters can be allcaps
	cl := false
	for _, c := range remark {
		if unicode.IsLetter(c) {
			cl = true
		}
	}

	switch {
	case remark == "": //remarks is ""
		return "Fine. Be that way!"
	case cl && remark == strings.ToUpper(remark) && string(remark[len(remark)-1]) == "?": //all caps question
		return "Calm down, I know what I'm doing!"
	case string(remark[len(remark)-1]) == "?": //remarks ends in "?"
		return "Sure."
	case cl && remark == strings.ToUpper(remark): //remarks is all caps
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}

}
