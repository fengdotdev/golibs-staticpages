package validators

import "errors"

// returns nil if the css is valid, otherwise returns an error
func CSSValidator(css string) error {
	if CSSIsValid(css) {
		return nil
	}

	return errors.New("Invalid CSS")
}

// returns true if the css is valid, otherwise returns false
func CSSIsValid(css string) bool {
	panic("Not implemented")
	return true
}
