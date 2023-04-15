package pkg

import (
	"reflect"
	"testing"
)

func TestIsValidateEmail(t *testing.T) {
	testCases := []struct {
		testCase    string
		expected    bool
		description string
	}{
		{testCase: "cheems@gmial.com", expected: true, description: "valid case"},
		{testCase: "cheems@gmialcom", expected: false, description: "invalid case"},
		{testCase: "", expected: false, description: "empty case"},
	}

	var rs bool
	for _, test := range testCases {
		rs = IsValidateEmail(test.testCase)
		if rs != test.expected {
			t.Errorf("Error in %s case, expected %t and got %t", test.description, test.expected, rs)
		}

	}
}

func TestIsValidateEmailList(t *testing.T) {
	testCases := []struct {
		testCase       []string
		expValidList   []string
		expInvalidList []string
		description    string
	}{
		{
			testCase:       []string{"cheems@gmial.com", "bad-format"},
			expValidList:   []string{"cheems@gmial.com"},
			expInvalidList: []string{"bad-format"},
			description:    "ona valid mail and one invalid mail",
		},
	}

	var vl []string
	var ivl []string
	for _, test := range testCases {
		vl, ivl = IsValidateEmailList(test.testCase)
		if !reflect.DeepEqual(vl, test.expValidList) {
			t.Errorf("Error in %s case, expected %v and got %v", test.description, test.expValidList, vl)
		}

		if !reflect.DeepEqual(ivl, test.expInvalidList) {
			t.Errorf("Error in %s case, expected %v and got %v", test.description, test.expInvalidList, ivl)
		}
	}
}

func TestIsValidSMTPServers(t *testing.T) {
	testCases := []struct {
		testCase    string
		expected    bool
		description string
	}{
		{testCase: "smtp.gmail.com", expected: true, description: "valid"},
		{testCase: "smtp.mail.yahoo.com", expected: true, description: "valid with subdomains"},
		{testCase: "ftp.", expected: false, description: "invalid"},
	}

	var rs bool
	for _, test := range testCases {
		rs = IsValidSMTPServer(test.testCase)
		if rs != test.expected {
			t.Errorf("Error in %s case, expected %t and got %t", test.description, test.expected, rs)
		}

	}
}
