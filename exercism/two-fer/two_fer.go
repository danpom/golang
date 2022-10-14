// Package twofer Given a name, return a string with the message:
// "One for <name>, one for me."
// if no name return "One for you, one for me."

package twofer

// ShareWith takes in a name and returns "One for <name>, one for me." or if name is nil returns "One for you, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
