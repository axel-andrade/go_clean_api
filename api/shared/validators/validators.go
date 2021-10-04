package validators

import "regexp"

func IsValidEmail(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func IsValidPassword(p string) bool {
	length := len(p)
	return length >= 6
}

func IsEmpty(str string) bool {
	length := len(str)
	return length <= 0
}
