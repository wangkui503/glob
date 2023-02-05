package syntax

import (
	"github.com/gobwas/glob/syntax/ast"
	"github.com/gobwas/glob/syntax/lexer"
)

func Parse(s string, sep []rune) (*ast.Node, error) {
	return ast.Parse(lexer.NewLexer(s, sep))
}

func Special(b byte) bool {
	return lexer.Special(b)
}
