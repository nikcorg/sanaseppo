SHELL := bash
BINARY_NAME := seppo
WORDLIST_SOURCE := data/nykysuomensanalista2024.txt
WORDLIST := data/sanalista.txt

.PHONY: build wordlist

build:
	go build -o "bin/$(BINARY_NAME)" -ldflags "-s -X 'main.binaryName=$(BINARY_NAME)'" *.go

wordlist: export temp=$(mktemp)
wordlist:
	export temp="$$(mktemp)" && \
	grep -vE "[-‑]" <(tail -n +2 "$(WORDLIST_SOURCE)" | awk '{ print $$1 }') | grep -E "...." > "$$temp" && \
	uniq "$$temp" > "$(WORDLIST)"