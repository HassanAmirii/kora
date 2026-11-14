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
func isTwoCharOperators(ops string) bool {
	twoCharOps := map[string]bool{
		">=": true,
		"<=": true,
		"!=": true,
		"||": true,
		"==": true,
		"+=": true,
		"-=": true,
		"&&": true,
		"++": true,
		"--": true,
	}
	return twoCharOps[ops]
}
func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch == '_')
}

func isOperator(ch byte) bool {
	return ch == '+' || ch == '*' || ch == '-' || ch == '/' || ch == '=' || ch == '!' || ch == '>' || ch == '<'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isDelimeter(ch byte) bool {
	return ch == '(' || ch == ')' || ch == '{' || ch == '}'
}
func main() {
	source := `set age >= 21
set name = "kenny"`

	i := 0
	line := 1
	column := 1

	for i < len(source) {
		ch := source[i]
		if i+1 < len(source) && isTwoCharOperators(source[i:i+2]) {
			printToken(Token{Type: "OPERATOR", Lexeme: source[i : i+2], Line: line, Column: column})
			i += 2
			column++
		} else if ch == '"' {
			start := i
			i++ // Skip the opening quote

			for i < len(source) && source[i] != '"' {
				i++
				column++
			}

			if i >= len(source) {
				fmt.Println("Error: Unterminated string")
			} else {
				lexeme := source[start+1 : i] // Grab text inside the quotes
				i++                           // Skip the closing quote
				printToken(Token{
					Type:   "STRING",
					Lexeme: lexeme,
					Line:   line,
					Column: column,
				})
			}
		} else if ch == '\n' {
			line++
			column = 1
			i++
		} else if ch == ' ' || ch == '\t' {
			i++
			column++
		} else if isLetter(ch) {
			start := i
			for i < len(source) && isLetter(source[i]) {
				i++
				column++
			}
			lexeme := source[start:i]
			printToken(Token{Type: wordType(lexeme), Lexeme: lexeme, Line: line, Column: column})
		} else if isDigit(ch) {
			start := i
			for i < len(source) && isDigit(source[i]) {
				i++
				column++
			}
			lexeme := source[start:i]
			printToken(Token{Type: "NUMBER", Lexeme: lexeme, Line: line, Column: column})
		} else if isOperator(ch) {
			printToken(Token{Type: "OPERATOR", Lexeme: string(ch), Line: line, Column: column})
			i++
			column++

		} else if isDelimeter(ch) {

			printToken(Token{Type: "DELIMETER", Lexeme: string(ch), Line: line, Column: column})
			i++

		} else {
			fmt.Println("Unexpected character '%c' at '%d':'%d' ", ch, 1, i+1)
		}
	}
}
