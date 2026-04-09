package alphabet

import (
	"strings"
	"testing"
)

func TestMatches(t *testing.T) {
	tcs := []struct {
		alphabet, match string
		shouldMatch     bool
	}{
		{"hspnuoe", "puhe", true},
		{"hspnuoe", "nuoli", false},
	}

	for _, tc := range tcs {
		s := strings.Split(tc.alphabet, "")
		a := New(s[0], s[1:])

		if matched := a.Matches(tc.match); matched != tc.shouldMatch {
			t.Errorf("expected %v for %q with alphabet %q, but got %v",
				tc.shouldMatch, tc.match, tc.alphabet, matched)
			t.Fail()
		}
	}
}

func TestPangram(t *testing.T) {
	tcs := []struct {
		alphabet, match  string
		pangram, perfect bool
	}{
		{"hspnuoe", "puhe", false, false},
		{"hspnuoe", "puhenopeus", true, false},
		{"hspnuoe", "huhupuhe", false, false},
		{"hspnuoe", "hspnuoe", true, true},
	}

	for _, tc := range tcs {
		s := strings.Split(tc.alphabet, "")
		a := New(s[0], s[1:])

		if pangram, perfect := a.IsPangram(tc.match); pangram != tc.pangram || perfect != tc.perfect {
			t.Errorf("expected (%v, %v) for %q with alphabet %q, but got (%v, %v)",
				tc.pangram, tc.perfect, tc.match, tc.alphabet, pangram, perfect)
			t.Fail()
		}
	}
}
