package utils

import (
	"errors"
	"fmt"
	"strings"
)

type DisplayOptions struct {
	Delimiter string
	LineEnd   string
}

func StringDisplay(stringsToDisplay []string, opts *DisplayOptions) (string, error) {
	if len(stringsToDisplay) == 0 {
		return "", errors.New("expected at least one string input - got an empty array")
	}

	// Set display options, defaulting if not provided
	defaultDisplayOpts := DisplayOptions{Delimiter: ", ", LineEnd: "."}
	delimiter := defaultDisplayOpts.Delimiter
	lineEnd := defaultDisplayOpts.LineEnd
	if opts != nil && opts.Delimiter != "" {
		delimiter = opts.Delimiter
	}
	if opts != nil && opts.LineEnd != "" {
		lineEnd = opts.LineEnd
	}

	concatenated := strings.Join(stringsToDisplay, delimiter)
	return fmt.Sprintf("%s%s", concatenated, lineEnd), nil
}
