package reverse

func Reverse(input string) string {
	var output string

	for _, v := range input {
		output = string(v) + output
	}
	return output
}
