package helpers

import "regexp"

func Regex(text string, expr string, replace string) string {
	reg, err := regexp.Compile(expr)
	StayCalm(err)

	return reg.ReplaceAllString(text, replace)
}

func OnlyNumber(text string) string {
	return Regex(text, "[^0-9]+", "")
}

func OnlyString(text string) string {
	return Regex(text, "[^a-zA-Z]+", "")
}

func OnlyStringAndSpace(text string) string {
	return Regex(text, "[^a-zA-Z ]+", "")
}

func OnlyAlphaNumeric(text string) string {
	return Regex(text, "[^a-zA-Z0-9]+", "")
}
