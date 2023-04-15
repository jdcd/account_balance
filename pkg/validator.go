package pkg

import (
	"regexp"
	"strconv"
)

func IsValidateEmailList(emailList []string) ([]string, []string) {
	validList := make([]string, 0)
	invalidList := make([]string, 0)

	for _, email := range emailList {
		if IsValidateEmail(email) {
			validList = append(validList, email)
		} else {
			invalidList = append(invalidList, email)
		}
	}

	return validList, invalidList

}

func IsValidateEmail(email string) bool {
	validEmailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(validEmailPattern)

	return regex.MatchString(email)
}

func IsValidStringOfNumbers(text string) bool {
	_, err := strconv.Atoi(text)
	return err == nil
}

func IsValidSMTPServer(text string) bool {
	regex := regexp.MustCompile(`^([a-zA-Z0-9.-]+\.[a-zA-Z]{2,})$`)
	return regex.MatchString(text)
}
