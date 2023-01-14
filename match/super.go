package match

import (
	"fmt"

	"github.com/gobwas/glob/util/strings"
)

type Super struct {
	Raw        string
	Separators []rune
	Sensitive  bool
}

func NewSuper(raw string, sep []rune) Super {
	return Super{raw, sep, strings.IndexAnyRunes(raw, sep) > 0}
}

func (self Super) Match(s string) bool {
	if self.Sensitive && len(s) > 0 && self.Raw[len(self.Raw)-1] != s[len(s)-1] {
		return false
	}
	return true
}

func (self Super) Len() int {
	return lenNo
}

func (self Super) Index(s string) (int, []int) {
	segments := acquireSegments(len(s) + 1)
	for i := range s {
		segments = append(segments, i)
	}
	segments = append(segments, len(s))

	return 0, segments
}

func (self Super) String() string {
	return fmt.Sprintf("<super:![%s]%s>", string(self.Separators), self.Raw)
}
