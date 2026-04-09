package dict

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"sort"
	"strings"
	"unicode/utf8"
)

const prefixSize = 2

// WordMatcher is for filtering the dictionary input.
type WordMatcher interface {
	Matches(w string) bool
}

// Dict represents a set of words matching a filter
type Dict struct {
	words        []string
	prefixes     map[string]int
	prefixCounts map[string]map[int]int
}

// New returns a pointed `Dict`, initialised to the wordlist input filtered using the `WordMatcher`.
func New(wordlist io.Reader, wm WordMatcher) (*Dict, error) {
	d := Dict{
		words:        []string{},
		prefixes:     map[string]int{},
		prefixCounts: map[string]map[int]int{},
	}

	err := d.readDict(wordlist, wm)

	return &d, err
}

func (d Dict) String() string {
	return fmt.Sprintf("Dict{size: %d, prefixes: %s}", len(d.words), strings.Join(getSortedKeys(d.prefixes), ", "))
}

// Size returns the number of words in the dictionary.
func (d Dict) Size() int {
	return len(d.words)
}

// MatchedWordes returns all matched words as an unsorted slice.
func (d Dict) MatchedWords() []string {
	return d.words
}

// Prefixes returns the prefixes for the matched words as a sorted slice.
func (d Dict) Prefixes() []string {
	return getSortedKeys(d.prefixes)
}

// PrefixMatches returns the number of matched words for the prefix.
func (d Dict) PrefixMatches(pfx string) int {
	return d.prefixes[pfx]
}

// WordLengths returns a slice of `{a, b int}` "tuples", where `a` is the word length and `b` is
// the number of matched words for the prefix.
func (d Dict) WordLengths(pfx string) [][2]int {
	counts := [][2]int{}

	for _, l := range getSortedKeys(d.prefixCounts[pfx]) {
		counts = append(counts, [2]int{l, d.prefixCounts[pfx][l]})
	}

	return counts
}

func (d *Dict) readDict(wordlist io.Reader, m WordMatcher) error {
	s := bufio.NewScanner(wordlist)

	for s.Scan() {
		ws := s.Text()

		if ws == "" || !m.Matches(ws) {
			continue
		}

		d.append(ws)
	}

	return s.Err()
}

func (d *Dict) append(w string) {
	d.words = append(d.words, w)

	wpfx := getPrefix(w, prefixSize)

	d.prefixes[wpfx]++

	if _, ok := d.prefixCounts[wpfx]; !ok {
		d.prefixCounts[wpfx] = map[int]int{}
	}

	runeCount := utf8.RuneCountInString(w)
	d.prefixCounts[wpfx][runeCount]++
}

func getPrefix(s string, size int) string {
	var (
		out       strings.Builder
		runeCount = 0
	)

	for _, r := range s {
		out.WriteRune(r)
		if runeCount++; runeCount == size {
			break
		}
	}

	return out.String()
}

func getSortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	ks := []K{}

	for k := range m {
		ks = append(ks, k)
	}

	sort.SliceStable(ks, func(a, b int) bool {
		return ks[a] < ks[b]
	})

	return ks
}
