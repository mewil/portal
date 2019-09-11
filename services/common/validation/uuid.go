package validation

import "github.com/google/uuid"

func ValidUUID(s string) error {
	_, err := uuid.Parse(s)
	return err
}
