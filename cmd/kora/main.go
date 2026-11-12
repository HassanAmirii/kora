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

func main() {
	t1 := Token{Type: "KEYWORD", Lexeme: "let", Line: 1, Column: 1}
	t2 := Token{Type: "IDENTIFIER", Lexeme: "age", Line: 1, Column: 5}
	t3 := Token{Type: "ASSIGN", Lexeme: "=", Line: 1, Column: 9}
	t4 := Token{Type: "NUMBER", Lexeme: "21", Line: 1, Column: 11}
	t5 := Token{Type: "NUMBER"}
	printToken(t1)
	printToken(t2)
	printToken(t3)
	printToken(t4)
	printToken(t5)

	source := "let age = 21"
	for i := 0; i < len(source); i++ {
		ch := source[i]
		if ch == ' '{
			continue
		}
		fmt.Println(i+1, string(ch))
	}
}
