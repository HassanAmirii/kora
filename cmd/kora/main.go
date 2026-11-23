package main

import (
	"fmt"
	"strconv"
)

type Token struct {
	Type   string
	Lexeme string
	Line   int
	Column int
}

type SetStatement struct {
	Name  string
	Value int
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

func isDelimiter(ch byte) bool {
	return ch == '(' || ch == ')' || ch == '{' || ch == '}'
}

func tokenize(source string) []Token {
	tokens := []Token{}

	i := 0
	line := 1
	lineStart := 0

	for i < len(source) {
		ch := source[i]
		column := i - lineStart + 1
		if i+1 < len(source) && isTwoCharOperators(source[i:i+2]) {

			tokens = append(tokens, Token{Type: "OPERATOR", Lexeme: source[i : i+2], Line: line, Column: column})
			i += 2
		} else if ch == '"' {
			start := i
			i++ // Skip the opening quote

			for i < len(source) && source[i] != '"' {
				i++
			}

			if i >= len(source) {
				fmt.Println("Error: Unterminated string")
			} else {
				lexeme := source[start+1 : i] // Grab text inside the quotes
				i++                           // Skip the closing quote
				tokens = append(tokens, Token{
					Type:   "STRING",
					Lexeme: lexeme,
					Line:   line,
					Column: column,
				})
			}
		} else if ch == '\n' {

			tokens = append(tokens, Token{Type: "NEWLINE", Lexeme: string(ch), Line: line, Column: column})
			line++
			lineStart = i + 1
			i++
		} else if ch == ' ' || ch == '\t' {
			i++
		} else if isLetter(ch) {
			start := i
			for i < len(source) && isLetter(source[i]) {
				i++
			}
			lexeme := source[start:i]
			tokens = append(tokens, Token{Type: wordType(lexeme), Lexeme: lexeme, Line: line, Column: column})
		} else if isDigit(ch) {
			start := i
			for i < len(source) && isDigit(source[i]) {
				i++
			}
			lexeme := source[start:i]
			tokens = append(tokens, Token{Type: "NUMBER", Lexeme: lexeme, Line: line, Column: column})
		} else if isOperator(ch) {

			tokens = append(tokens, Token{Type: "OPERATOR", Lexeme: string(ch), Line: line, Column: column})
			i++

		} else if isDelimiter(ch) {

			tokens = append(tokens, Token{Type: "DELIMITER", Lexeme: string(ch), Line: line, Column: column})
			i++

		} else {
			fmt.Printf("Unexpected character '%c' at %d:%d \n", ch, line, column)
			i++
		}
	}

	return tokens

}

func parseSetStatement(tokens []Token) (SetStatement, error) {

	if len(tokens) != 4 {
		return SetStatement{}, fmt.Errorf("expected 4 tokens, but found %d", len(tokens))
	}

	firstToken := tokens[0]
	secondToken := tokens[1]
	thirdToken := tokens[2]
	forthToken := tokens[3]

	if firstToken.Type != "KEYWORD" || firstToken.Lexeme != "set" {
		return SetStatement{}, fmt.Errorf("%d:%d expected the KEYWORD : set. but found %q: %q", firstToken.Line, firstToken.Column, firstToken.Type, firstToken.Lexeme)
	}

	if secondToken.Type != "IDENTIFIER" {
		return SetStatement{}, fmt.Errorf("%d:%d expected an IDENTIFIER. but found %q", secondToken.Line, secondToken.Column, secondToken.Type)
	}

	if thirdToken.Type != "OPERATOR" || thirdToken.Lexeme != "=" {
		return SetStatement{}, fmt.Errorf("%d:%d expected an OPERATOR (=). but found %q", thirdToken.Line, thirdToken.Column, thirdToken.Lexeme)
	}

	if forthToken.Type != "NUMBER" {
		return SetStatement{}, fmt.Errorf("%d:%d expected a NUMBER. but found %q", forthToken.Line, forthToken.Column, forthToken.Type)
	}
	value, err := strconv.Atoi(forthToken.Lexeme)

	if err != nil {
		return SetStatement{}, fmt.Errorf("failed to parse value: %w", err)
	}
	return SetStatement{
		Name:  secondToken.Lexeme,
		Value: value,
	}, nil

}

func main() {

	source := "love age + 21"

	tokens := tokenize(source)
	stmt, err := parseSetStatement(tokens)

	if err != nil {
		fmt.Println("parse error:", err)
		return
	}
	fmt.Println(stmt)

	for _, tokenValue := range tokens {
		fmt.Printf("Type: %s | Lexeme: %q | Line: %d, Column: %d \n", tokenValue.Type, tokenValue.Lexeme, tokenValue.Line, tokenValue.Column)
	}

}
