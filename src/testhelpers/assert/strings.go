package assert

import (
	"strings"
	"testing"
	"fmt"
)


func JSONStringEquals(t *testing.T, actual string, expected string, msgAndArgs ...interface{}) bool {
	if RemoveWhitespace(actual) == RemoveWhitespace(expected){
		return true
	}
	Fail(t, fmt.Sprintf("Not equal:\n%s\n%s", actual,expected), msgAndArgs...)
	return false
}

func RemoveWhitespace(body string) string {
	body = strings.Replace(body, " ", "", -1)
	body = strings.Replace(body, "\n", "", -1)
	body = strings.Replace(body, "\r", "", -1)
	body = strings.Replace(body, "\t", "", -1)
	return body
}

