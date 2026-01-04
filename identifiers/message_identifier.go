package identifiers

type MessageIdentifier struct {}

func (mi MessageIdentifier) Name() string {
	return "message"
}

func (mi MessageIdentifier) GetValue(vals *map[string]string) (string, error) {
	if val, ok := (*vals)[mi.Name()]; ok {
		return val, nil
	}
	return "", nil
}


