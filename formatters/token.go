package formatters

import (
	"strings"
	"github.com/hashen47/loggo/identifiers"
)

type Token struct {
	isIdentifier bool 
	identifier   identifiers.IdentifierI
	text         string
}

func (token Token) Compare(other Token) bool {
	if token.isIdentifier != other.isIdentifier {
		return false
	}

	if token.identifier != other.identifier {
		return false
	}

	if strings.Compare(token.text, other.text) != 0 {
		return false
	}

	return true
}

func Tokenizer(format string, registeredIdentifiers map[string]identifiers.IdentifierI) []Token {
	tokens := make([]Token, 0)

	if strings.Compare(strings.Trim(format, " "), "") == 0 {
		return tokens
	}

	textRunes 		    := make([]rune, 0)
	identifierNameRunes := make([]rune, 0)
	isIdentifierStart   := false

	for i, ch := range format {
		if ch == '%' {
			if isIdentifierStart {
				isIdentifierStart = false
				identifierName    := string(identifierNameRunes)
				isIdentifier      := false

				if _, ok := registeredIdentifiers[identifierName]; ok {
					isIdentifier = true
				}

				if isIdentifier {
					token := Token{true, registeredIdentifiers[identifierName], ""}
					tokens = append(tokens, token)
				} else {
					isIdentifierStart = true
					token := Token{false, identifiers.EmptyIdentifier{}, "%" + identifierName}
					if i == len(format)-1 {
						token.text += "%"
					}
					tokens = append(tokens, token)
				}

				identifierNameRunes = make([]rune, 0)
			} else {
				if len(textRunes) > 0 {
					token := Token{false, identifiers.EmptyIdentifier{}, string(textRunes)}
					tokens = append(tokens, token)
					textRunes = make([]rune, 0)
				}
				isIdentifierStart = true
				if i == len(format)-1 {
					token := Token{false, identifiers.EmptyIdentifier{}, "%"}
					tokens = append(tokens, token)
				}
			}
		} else {
			if isIdentifierStart {
				identifierNameRunes = append(identifierNameRunes, ch)
			} else {
				textRunes = append(textRunes, ch)
			}
		}
	}

	if isIdentifierStart {
		if len(identifierNameRunes) > 0 {
			identifierNameRunes = append([]rune{'%'}, identifierNameRunes...)
			token := Token{false, identifiers.EmptyIdentifier{}, string(identifierNameRunes)}
			tokens = append(tokens, token)
		}
	} else {
		if len(textRunes) > 0 {
			token := Token{false, identifiers.EmptyIdentifier{}, string(textRunes)}
			tokens = append(tokens, token)
		}
	}

	return tokens
}
