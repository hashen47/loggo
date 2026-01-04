package identifiers

import (
	"fmt"
)

type LogLevelIdentifier struct {}

func (lli LogLevelIdentifier) Name() string {
	return "log_level"
}

func (lli LogLevelIdentifier) GetValue(vals *map[string]string) (string, error) {
	if level, ok := (*vals)[lli.Name()]; ok {
		switch LogLevel(level) {
		case Debug:
			return "DEBUG", nil
		case Info:
			return "INFO", nil
		case Warn:
			return "WARN", nil
		case Error:
			return "ERROR", nil
		default:
			return "", &ErrInvalidIdentifierVal{msg: fmt.Sprintf("unknown %s value: %s", lli.Name(), level)}
		}
	}
	return "", nil
}

