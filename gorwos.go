package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type options struct {
	uppercase bool
	lowercase bool
	digits    bool
	symbols   bool
	prefix    bool
	suffix    bool
	length    int
	count     int
}

func main() {

	var length = flag.Int("l", 16, "length of string")
	var count = flag.Int("c", 1, "number of passwords")
	var customOptions = flag.String("o", "ulds", "custom options (u, l, d, s, p, x")
	flag.Parse()

	err := run(*count, *length, *customOptions)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(count, length int, customOptions string) error {
	defaultOptions := &options{
		uppercase: false,
		lowercase: false,
		digits:    false,
		symbols:   false,
		count:     count,
		length:    length,
	}

	if customOptions == "" {
		return errors.New("ERR: No options specified.")
	}

	if strings.Contains(customOptions, "u") {
		defaultOptions.uppercase = true
	}

	if strings.Contains(customOptions, "l") {
		defaultOptions.lowercase = true
	}

	if strings.Contains(customOptions, "d") {
		defaultOptions.digits = true
	}

	if strings.Contains(customOptions, "s") {
		defaultOptions.symbols = true
	}

	if strings.Contains(customOptions, "p") {
		defaultOptions.prefix = true
	}

	if strings.Contains(customOptions, "x") {
		defaultOptions.suffix = true
	}

	password, err := newRandomString(defaultOptions)
	if err != nil {
		return errors.New(err.Error())
	}

	for _, p := range password {
		fmt.Println(p)
	}

	return nil
}

func shuffleString(str string) []byte {
	rand.Seed(time.Now().Unix())

	data := []byte(str)
	rand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})

	return data
}

func generateKeyspace(opt *options) []byte {
	var keyspace []string

	if opt.uppercase {
		keyspace = append(keyspace, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}

	if opt.lowercase {
		keyspace = append(keyspace, "abcdefghijklmnopqrstuvwxyz")
	}

	if opt.digits {
		keyspace = append(keyspace, "0123456789")
	}

	if opt.symbols {
		keyspace = append(keyspace, "@!-*#%=")
	}

	keyspaceJoined := strings.Join(keyspace, "")

	return shuffleString(keyspaceJoined)
}

func newRandomString(opt *options) ([]string, error) {

	if !opt.uppercase && !opt.lowercase && !opt.digits && !opt.symbols {
		return nil, errors.New("ERR: You must specify atleast of one: uppercase, lowercase, digits or symbols.")
	}

	keyspace := generateKeyspace(opt)
	var passwords []string

	for i := 0; i < opt.count; i++ {
		var password []byte

		for j := 0; j < opt.length; j++ {
			password = append(password, keyspace[rand.Intn(len(keyspace)-1)])
		}

		passwords = append(passwords, string(password))
	}

	return passwords, nil
}
