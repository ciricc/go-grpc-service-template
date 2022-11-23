package printer

import "errors"

func validateMessage(msg string) error {
	if msg == "" {
		return errors.New("message mepty")
	}
	return nil
}
