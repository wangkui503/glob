package match

import (
	"fmt"
	gostrings "strings"

	"github.com/gobwas/glob/util/strings"
)

type Any struct {
	Raw        string
	Separators []rune
	SuffixSep  bool
	PrefixSep  bool
}

func NewAny(raw string, s []rune) Any {
	return Any{raw, s, strings.LastIndexAnyRunes(raw, s) == len(raw)-1, strings.IndexAnyRunes(raw, s) == 0}
}

func (self Any) Match(s string) bool {
	if len(s) < 1 {
		if self.SuffixSep && gostrings.HasPrefix(self.Raw, string(0)) ||
			self.PrefixSep && gostrings.HasSuffix(self.Raw, string(0)) {
			return false
		} else {
			return true
		}
	} else if (self.PrefixSep || gostrings.HasPrefix(self.Raw, string(0))) && gostrings.HasPrefix(s, ".") {
		return false
	}
	return strings.IndexAnyRunes(s, self.Separators) == -1
}

func (self Any) Index(s string) (int, []int) {
	if len(s) < 1 {
		return -1, nil
	}
	if (self.PrefixSep || gostrings.HasPrefix(self.Raw, string(0))) && gostrings.HasPrefix(s, ".") {
		return -1, nil
	}
	found := strings.IndexAnyRunes(s, self.Separators)
	switch found {
	case -1:
	case 0:
		return 0, segments0
	default:
		s = s[:found]
	}

	segments := acquireSegments(len(s))
	for i := range s {
		segments = append(segments, i)
	}
	segments = append(segments, len(s))

	return 0, segments
}

func (self Any) Len() int {
	return lenNo
}

func (self Any) String() string {
	return fmt.Sprintf("<any:![%s]%s>", string(self.Separators), self.Raw)
}
