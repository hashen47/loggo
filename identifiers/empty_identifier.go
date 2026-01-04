package identifiers

type EmptyIdentifier struct {}

func (ei EmptyIdentifier) Name() string {
	return ""
}

func (ei EmptyIdentifier) GetValue(vals *map[string]string) (string, error) {
	return "", nil
}
