package day13

import (
	"fmt"
	"io"
)

func Parse(source io.Reader) ([]Entry, error) {
	entries := make([]Entry, 0, 100)
	for {
		var entry Entry
		if err := parseEntry(source, &entry); err == nil {
			entries = append(entries, entry)
		} else {
			return entries, err
		}
	}
}

func parseEntry(source io.Reader, entry *Entry) error {
	var change string
	_, err := fmt.Fscanf(
		"%s would %s %d happiness units by sitting next to %s.\n",
		&entry.Person, &change, &entry.Happiness, &entry.Other)
	if err == nil && change == "lose" {
		entry.Happiness *= -1
	}
	return err
}
