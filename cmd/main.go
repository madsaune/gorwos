package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/madsaune/gorwos/pkg/gorwos"
)

var (
	length        = flag.Int("l", 16, "length of string")
	count         = flag.Int("c", 1, "number of passwords")
	customOptions = flag.String("o", "ulds", "custom options (u, l, d, s, p, x")
)

func main() {
	flag.Parse()

	opts := optionsFromArgs(*count, *length, *customOptions)

	passwords, err := gorwos.GenerateString(opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, v := range passwords {
		fmt.Fprintln(os.Stdout, v)
	}
}

func optionsFromArgs(count, length int, customOptions string) gorwos.Options {
	opts := gorwos.Options{
		Length: length,
		Count:  count,
	}

	if strings.Contains(customOptions, "u") {
		opts.Uppercase = true
	}

	if strings.Contains(customOptions, "l") {
		opts.Lowercase = true
	}

	if strings.Contains(customOptions, "d") {
		opts.Digits = true
	}

	if strings.Contains(customOptions, "s") {
		opts.Symbols = true
	}

	if strings.Contains(customOptions, "p") {
		opts.Prefix = true
	}

	if strings.Contains(customOptions, "x") {
		opts.Suffix = true
	}

	return opts
}
