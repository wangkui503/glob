package glob

import (
	"github.com/gobwas/glob/compiler"
	"github.com/gobwas/glob/syntax"
	"github.com/gobwas/glob/syntax/ast"
)

// Glob represents compiled glob pattern.
type Glob interface {
	Match(string) bool
}

// Compile creates Glob for given pattern and strings (if any present after pattern) as separators.
// The pattern syntax is:
//
//    pattern:
//        { term }
//
//    term:
//        `*`         matches any sequence of non-separator characters
//        `**`        matches any sequence of characters
//        `?`         matches any single non-separator character
//        `[` [ `!` ] { character-range } `]`
//                    character class (must be non-empty)
//        `{` pattern-list `}`
//                    pattern alternatives
//        c           matches character c (c != `*`, `**`, `?`, `\`, `[`, `{`, `}`)
//        `\` c       matches character c
//
//    character-range:
//        c           matches character c (c != `\\`, `-`, `]`)
//        `\` c       matches character c
//        lo `-` hi   matches character c for lo <= c <= hi
//
//    pattern-list:
//        pattern { `,` pattern }
//                    comma-separated (without spaces) patterns
//
func Compile(pattern string, separators ...rune) (Glob, error) {
	ast, err := syntax.Parse(pattern, separators)
	if err != nil {
		return nil, err
	}
	if !hasKind(ast) {
		return nil, nil
	}

	matcher, err := compiler.Compile(ast, separators)
	if err != nil {
		return nil, err
	}

	return matcher, nil
}

func IsGlob(name string, separators ...rune) (bool, error) {
	node, err := syntax.Parse(name, separators)
	if err != nil {
		return false, err
	}
	return hasKind(node), nil
}

func hasKind(node *ast.Node) bool {
	/* if node == nil {
		return false
	} */
	switch node.Kind {
	case ast.KindList, ast.KindRange, ast.KindAny, ast.KindSuper, ast.KindSingle, ast.KindAnyOf:
		return true
	}
	for _, node = range node.Children {
		if hasKind(node) {
			return true
		}
	}
	return false
}

// MustCompile is the same as Compile, except that if Compile returns error, this will panic
func MustCompile(pattern string, separators ...rune) Glob {
	g, err := Compile(pattern, separators...)
	if err != nil {
		panic(err)
	}

	return g
}

// QuoteMeta returns a string that quotes all glob pattern meta characters
// inside the argument text; For example, QuoteMeta(`{foo*}`) returns `\[foo\*\]`.
func QuoteMeta(s string) string {
	b := make([]byte, 2*len(s))

	// a byte loop is correct because all meta characters are ASCII
	j := 0
	for i := 0; i < len(s); i++ {
		if syntax.Special(s[i]) {
			b[j] = '\\'
			j++
		}
		b[j] = s[i]
		j++
	}

	return string(b[0:j])
}
