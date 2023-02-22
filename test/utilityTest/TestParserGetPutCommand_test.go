package test

import (
	"testing"

	"rest.com.tw/tinymud/src/RestGo.Util/utility"
)

func TestParserGetPutCommand(t *testing.T) {
	testCases := []struct {
		msg           string
		objectNum     int
		objectName    string
		containNum    int
		containerName string
	}{
		{"3.book into 4.shelf", 3, "book", 4, "shelf"},
		{"3.book in 4.shelf", 3, "book", 4, "shelf"},
		{"3.book 4.shelf", 3, "book", 4, "shelf"},
		{"book into shelf", 1, "book", 1, "shelf"},
		{"book in shelf", 1, "book", 1, "shelf"},
		{"book shelf", 1, "book", 1, "shelf"},
	}
	for _, tc := range testCases {
		objectNum, objectName, containNum, containerName := utility.ParserGetPutCommand(tc.msg)
		if objectNum != tc.objectNum {
			t.Errorf("For message '%s', expected objectNum to be %d but got %d", tc.msg, tc.objectNum, objectNum)
		}
		if objectName != tc.objectName {
			t.Errorf("For message '%s', expected objectName to be %s but got %s", tc.msg, tc.objectName, objectName)
		}
		if containNum != tc.containNum {
			t.Errorf("For message '%s', expected containNum to be %d but got %d", tc.msg, tc.containNum, containNum)
		}
		if containerName != tc.containerName {
			t.Errorf("For message '%s', expected containerName to be %s but got %s", tc.msg, tc.containerName, containerName)
		}
	}

}

//todo: 要寫utility.AlignMessageByBrackets()的unit test
