package match

import (
	"fmt"
	gostrings "strings"

	"github.com/gobwas/glob/util/strings"
)

type Super struct {
	Raw        string
	Separators []rune
	SuffixSep  bool
	PrefixSep  bool
}

func NewSuper(raw string, sep []rune) Super {
	return Super{raw, sep, strings.IndexAnyRunes(raw, sep) > 0, (strings.IndexAnyRunes(raw, sep) == 0 || gostrings.HasPrefix(raw, string(0)))}
}

func (self Super) containsSepDot(s string) (yes bool) {
	for len(s) > 0 {
		hasSep := strings.IndexAnyRunes(s, self.Separators)
		if hasSep < 0 {
			break
		}
		if s = s[hasSep+1:]; len(s) > 0 && s[0] == '.' {
			yes = true
			break
		}
	}
	return
}

func (self Super) Match(s string) bool {
	if self.SuffixSep && len(s) > 0 && self.Raw[len(self.Raw)-1] != s[len(s)-1] ||
		self.PrefixSep && gostrings.HasPrefix(s, ".") ||
		self.containsSepDot(s) {
		return false
	}
	return true
}

func (self Super) Len() int {
	return lenNo
}

func (self Super) Index(s string) (int, []int) {
	if self.PrefixSep && gostrings.HasPrefix(s, ".") ||
		self.containsSepDot(s) {
		return -1, nil
	}
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
