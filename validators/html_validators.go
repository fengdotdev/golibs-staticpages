package validators

import (
	"errors"
	"regexp"
)


const HTMLRegex = `(?i)^<([a-z]+)(\s[^<>]*)?>.*?</[a-z]+>$`

func HTMLisValid(html string) bool {

	re := regexp.MustCompile(HTMLRegex)
	return re.MatchString(html)
}

func HTMLValidator(html string) error {
	if HTMLisValid(html) {
		return nil
	}
	return errors.New("invalid HTML")
}
