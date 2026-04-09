package main

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"sanaseppo/alphabet"
	"sanaseppo/dict"
	"sanaseppo/logx"
)

const envOverrideDict = "SANASEPPO_DICT"

var (
	errMissingArguments = errors.New("expected at least two arguments")
	errReadingDict      = errors.New("error reading dictionary")
	binaryName          string
	logger              = logx.New()

	//go:embed data/sanalista.txt
	wordList string
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		switch {
		case errors.Is(errMissingArguments, err):
			fmt.Fprintf(os.Stderr, "fatal: %s\n", err)
			fmt.Fprintf(os.Stderr, "usage: %s <center letter> <alphabet> [<reveal prefix>]\n", binaryName)

		case errors.Is(errReadingDict, err):
			fmt.Fprintf(os.Stderr, "fatal: error reading wordlist: %s\n", err)

		default:
			fmt.Fprintf(os.Stderr, "fatal: unhandled error: %s\n", err)
		}

		os.Exit(1)
	}

	os.Exit(0)
}

func run(argv []string) error {
	if argc := len(argv); argc < 2 {
		return errMissingArguments
	}

	al := alphabet.New(argv[0], strings.Split(argv[1], ""))
	d, err := initDict(al)

	logger.Debug("init", "alphabet", al, "dictionary", d)

	if err != nil {
		return errReadingDict
	}

	reveal := false
	revealPrefix := ""
	if len(argv) > 2 {
		reveal = true
		revealPrefix = argv[2]
	}

	for _, p := range d.Prefixes() {
		ns := d.WordLengths(p)

		fmt.Printf("%s: %2d", p, d.PrefixMatches(p))

		for _, l := range ns {
			fmt.Printf("\t%d=%d", l[0], l[1])
		}

		fmt.Println()
	}

	if reveal {
		fmt.Println()
		fmt.Printf("revealing words for prefix %q\n", revealPrefix)
	}

	pgs, ppgs := 0, 0
	for _, w := range d.MatchedWords() {
		pangram, perfect := al.IsPangram(w)
		if pangram {
			pgs++
			if perfect {
				ppgs++
			}
		}

		if reveal && strings.HasPrefix(w, revealPrefix) {
			fmt.Printf("- %s", w)
			if pangram {
				fmt.Print(" *")
			}
			fmt.Println()
		}
	}

	fmt.Println()
	fmt.Printf("%d words (%d pangrams, %d perfect)\n", d.Size(), pgs, ppgs)

	return nil
}

func initDict(a alphabet.Alphabet) (*dict.Dict, error) {
	var r io.Reader = strings.NewReader(wordList)

	if wl, ok := os.LookupEnv(envOverrideDict); ok {
		logger.Debug("wordlist override", "source", wl)

		f, err := os.Open(wl)
		if err != nil {
			return nil, err
		}

		r = f
	}

	d, err := dict.New(r, a)

	return &d, err
}
