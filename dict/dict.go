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

type WordMatcher interface {
	Matches(w string) bool
}

type Dict struct {
	words        []string
	prefixes     map[string]int
	prefixCounts map[string]map[int]int
}

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

func (d Dict) Size() int {
	return len(d.words)
}

func (d Dict) MatchedWords() []string {
	return d.words
}

func (d Dict) Prefixes() []string {
	return getSortedKeys(d.prefixes)
}

func (d Dict) PrefixMatches(pfx string) int {
	return d.prefixes[pfx]
}

func (d Dict) WordLengths(pfx string) [][]int {
	counts := [][]int{}

	for _, l := range getSortedKeys(d.prefixCounts[pfx]) {
		counts = append(counts, []int{l, d.prefixCounts[pfx][l]})
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

	runeCount := utf8.RuneCountInString(w)

	if _, ok := d.prefixCounts[wpfx]; !ok {
		d.prefixCounts[wpfx] = map[int]int{}
	}

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
