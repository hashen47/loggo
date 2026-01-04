package identifiers

type IdentifierI interface {
	Name() string
	GetValue(vals *map[string]string) (string, error)
}

type LogLevel string
const (
	Debug LogLevel = "DEBUG"
	Info  LogLevel = "INFO"
	Warn  LogLevel = "WARN"
	Error LogLevel = "ERROR"
)
