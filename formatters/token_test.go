package formatters

import (
	"fmt"
	"testing"
	"github.com/hashen47/loggo/identifiers"
)

func TestTokenizer(t *testing.T) {
	idfs := []identifiers.IdentifierI{
		identifiers.DateTimeIdentifier{},
		identifiers.MessageIdentifier{},
		identifiers.LogLevelIdentifier{},
	}

	registeredIdentifiers := make(map[string]identifiers.IdentifierI, 0)

	for _, idf := range idfs {
		registeredIdentifiers[idf.Name()] = idf
	}

	type Testcase struct {
		format string
		tokens []Token
	}

	testcases := []Testcase{
		{
			format: fmt.Sprintf("[%%%s%%] [%%%s%%]: %%%s%%", identifiers.DateTimeIdentifier{}.Name(), identifiers.LogLevelIdentifier{}.Name(), identifiers.MessageIdentifier{}.Name()),
			tokens: []Token{
				{false, identifiers.EmptyIdentifier{}   , "["},
				{true , identifiers.DateTimeIdentifier{}, ""},
				{false, identifiers.EmptyIdentifier{}   , "] ["},
				{true , identifiers.LogLevelIdentifier{}, ""},
				{false, identifiers.EmptyIdentifier{}   , "]: "},
				{true , identifiers.MessageIdentifier{}, ""},
			},
		},
		{
			format: fmt.Sprintf("[%%%s%%]: %%%s%%", identifiers.LogLevelIdentifier{}.Name(), identifiers.MessageIdentifier{}.Name()),
			tokens: []Token{
				{false, identifiers.EmptyIdentifier{}   , "["},
				{true , identifiers.LogLevelIdentifier{}, ""},
				{false, identifiers.EmptyIdentifier{}   , "]: "},
				{true , identifiers.MessageIdentifier{}, ""},
			},
		},
		{
			format: fmt.Sprintf("[%%unregistered_identifier%%]: %%%s%%", identifiers.MessageIdentifier{}.Name()),
			tokens: []Token{
				{false, identifiers.EmptyIdentifier{}   , "["},
				{false, identifiers.EmptyIdentifier{}   , "%unregistered_identifier"},
				{false, identifiers.EmptyIdentifier{}   , "%]: "},
				{true , identifiers.MessageIdentifier{}, ""},
			},
		},
		{
			format: fmt.Sprintf("%%nothing%%%s%%%%else%%%%%s%%", identifiers.DateTimeIdentifier{}.Name(), identifiers.MessageIdentifier{}.Name()),
			tokens: []Token{
				{false, identifiers.EmptyIdentifier{}   , "%nothing"},
				{true , identifiers.DateTimeIdentifier{}, ""},
				{false, identifiers.EmptyIdentifier{}   , "%else"},
				{false, identifiers.EmptyIdentifier{}   , "%"},
				{true , identifiers.MessageIdentifier{} , ""},
			},
		},
		{
			format: "%nothing here%something%",
			tokens: []Token{
				{false, identifiers.EmptyIdentifier{}, "%nothing here"},
				{false, identifiers.EmptyIdentifier{}, "%something%"},
			},
		},
		{
			format: "%nothing here%something%t",
			tokens: []Token{
				{false, identifiers.EmptyIdentifier{}, "%nothing here"},
				{false, identifiers.EmptyIdentifier{}, "%something"},
				{false, identifiers.EmptyIdentifier{}, "%t"},
			},
		},
		{
			format: "%nothing here%something%%",
			tokens: []Token{
				{false, identifiers.EmptyIdentifier{}, "%nothing here"},
				{false, identifiers.EmptyIdentifier{}, "%something"},
				{false, identifiers.EmptyIdentifier{}, "%%"},
			},
		},
		{
			format: "no identifiers here (:",
			tokens: []Token{
				{false, identifiers.EmptyIdentifier{}, "no identifiers here (:"},
			},
		},
		{
			format: "no identifiers here (:%",
			tokens: []Token{
				{false, identifiers.EmptyIdentifier{}, "no identifiers here (:"},
				{false, identifiers.EmptyIdentifier{}, "%"},
			},
		},
		{
			format: "no identifiers %here (: ",
			tokens: []Token{
				{false, identifiers.EmptyIdentifier{}, "no identifiers "},
				{false, identifiers.EmptyIdentifier{}, "%here (: "},
			},
		},
		{
			format: "",
			tokens: []Token{},
		},
		{
			format: "                  ",
			tokens: []Token{},
		},
	}

	for _, tc := range testcases {
		tokens := Tokenizer(tc.format, registeredIdentifiers)
		if len(tokens) != len(tc.tokens) {
			t.Log(tokens);
			t.Fatalf("FAIL: format: %s, expect token length: %d, real token length: %d\n", tc.format, len(tc.tokens), len(tokens))
		}

		for i := 0; i < len(tokens); i++ {
			if !tc.tokens[i].Compare(tokens[i]) {
				t.Fatalf("FAIL: format: %s, expect token: %v, real token: %v\n", tc.format, tc.tokens[i], tokens[i])
			}
		}
	}
}
