package formatters

import (
	"github.com/hashen47/loggo/identifiers"
)

type Formatter struct {
	tokens *[]Token
}

func NewFormatter(formatText string, idfs ...identifiers.IdentifierI) *Formatter {
	f := Formatter{}
	registeredIdentifiers := make(map[string]identifiers.IdentifierI, 0)
	for _, idf := range idfs {
		registeredIdentifiers[idf.Name()] = idf
	}
	tokens := Tokenizer(formatText, registeredIdentifiers)
	f.tokens = &tokens
	return &f
}

func (f *Formatter) GetText(vals *map[string]string) (string, error) {
	output := ""
	for _, token := range *(f.tokens) {
		if token.isIdentifier {
			val, err := token.identifier.GetValue(vals)
			if err != nil {
				return "", err
			}
			output += val
		} else {
			output += token.text
		}
	}
	return output, nil
}
