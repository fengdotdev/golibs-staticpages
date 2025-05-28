package validators

import "errors"

func JSValidator(js string) error {
	if JSIsValid(js) {
		return nil
	}
	return errors.New("Invalid JS")
}

func JSIsValid(js string) bool {
	panic("Not implemented")
	return true
}
