package utility

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/logrusorgru/aurora/v4"
)

func ToCyan(str string, num ...int) string {
	upplen := 1
	if len(num) > 0 && num[0] > 0 {
		upplen = num[0]
	}

	colored := aurora.Cyan(str[:upplen])
	rtn := fmt.Sprintf("%s%s", colored, str[upplen:])
	return rtn
}

func ValidateStringLengRange(input string, min int, max int) error {
	reg, _ := regexp.Compile(fmt.Sprintf("^[a-zA-Z0-9 ]{%d,%d}$", min, max))
	if reg.MatchString(input) {
		return nil
	} else {
		return errors.New(aurora.Sprintf("請輸入%d~%d個以內的字元。", aurora.Cyan(min), aurora.Cyan(max)))
	}
}
