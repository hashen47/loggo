package identifiers

import (
	"fmt"
)

type ErrInvalidIdentifierVal struct {
	msg string
}

func (e *ErrInvalidIdentifierVal) Error() string {
	return fmt.Sprintf("Err: INVALID_IDENTIFIER_ERROR: %s", e.msg)
}
