package alphabet

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Alphabet struct {
	size    int
	letters []string
	centre  string
	excl    *regexp.Regexp
}

func New(centre string, letters []string) Alphabet {
	return Alphabet{
		size:    len(letters) + 1,
		letters: letters,
		centre:  centre,
		excl:    regexp.MustCompile(fmt.Sprintf("[^%s%s]", centre, strings.Join(letters, ""))),
	}
}

func (a Alphabet) String() string {
	return fmt.Sprintf("Alphabet{centre: %s, letters: %s}", a.centre, strings.Join(a.letters, ""))
}

func (a Alphabet) Matches(word string) bool {
	return !a.excl.MatchString(word) && strings.Contains(word, a.centre)
}

func (a Alphabet) IsPangram(word string) (bool, bool) {
	rc := utf8.RuneCountInString(word)

	if rc < a.size {
		return false, false
	}

	cs := map[rune]struct{}{}

	for _, r := range word {
		cs[r] = struct{}{}
	}

	return len(cs) == a.size, rc == a.size
}
