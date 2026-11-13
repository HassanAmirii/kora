package main

import "fmt"

type Token struct {
	Type   string
	Lexeme string
	Line   int
	Column int
}

func printToken(t Token) {
	fmt.Println("Type:", t.Type, "| Lexeme:", t.Lexeme, "| Line:", t.Line, "| Column:", t.Column)
}

func wordType(word string) string {

	keywords := map[string]bool{
		"set":    true,
		"if":     true,
		"else":   true,
		"output": true,
		"true":   true,
		"false":  true,
	}
	if keywords[word] {
		return "KEYWORD"
	}
	return "IDENTIFIER"

}
func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch == '_')
}

func main() {
	source := "set age = 21"
	temporary := []byte{}
	start := -1

	for i := 0; i < len(source); i++ {
		ch := source[i]
		if isLetter(ch) {
			if len(temporary) == 0 {
				start = i
			}
			temporary = append(temporary, ch)
		} else {
			if len(temporary) > 0 {
				lexeme := string(temporary)
				token := Token{
					Type:   wordType(lexeme),
					Lexeme: lexeme,
					Line:   1,
					Column: start + 1,
				}
				printToken(token)
				temporary = []byte{}
				start = -1
			}
		}
	}
}
