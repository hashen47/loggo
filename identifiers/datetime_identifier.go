package identifiers

import (
	"time"
)

type DateTimeIdentifier struct {}

func (dti DateTimeIdentifier) Name() string {
	return "date_time"
}

func (dti DateTimeIdentifier) GetValue(vals *map[string]string) (string, error) {
	return time.Now().Format(time.DateTime), nil
}
