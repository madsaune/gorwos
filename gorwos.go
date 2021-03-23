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
		return nil, errors.New("You must specify atleast of one: uppercase, lowercase, digits or symbols.")
	}

	keyspace := generateKeyspace(opt)
	var passwords []string
	var password []byte

	for i := 0; i < opt.length; i++ {
		password = append(password, keyspace[rand.Intn(len(keyspace)-1)])
	}

	passwords = append(passwords, string(password))

	return passwords, nil
}

func main() {

	var length = flag.Int("l", 16, "length of string")
	var count = flag.Int("c", 1, "number of passwords")
	var customOptions = flag.String("o", "ulds", "custom options (u, l, d, s, p, x")
	var passwordType = flag.String("t", "string", "type of password")
	flag.Parse()

	defaultOptions := &options{
		uppercase: false,
		lowercase: false,
		digits:    false,
		symbols:   false,
		length:    *length,
		count:     *count,
	}

	if *customOptions == "" {
		fmt.Fprintln(os.Stderr, "ERR: No options specified.")
		os.Exit(1)
	}

	if strings.Contains(*customOptions, "u") {
		defaultOptions.uppercase = true
	}

	if strings.Contains(*customOptions, "l") {
		defaultOptions.lowercase = true
	}

	if strings.Contains(*customOptions, "d") {
		defaultOptions.digits = true
	}

	if strings.Contains(*customOptions, "s") {
		defaultOptions.symbols = true
	}

	if strings.Contains(*customOptions, "p") {
		defaultOptions.prefix = true
	}

	if strings.Contains(*customOptions, "x") {
		defaultOptions.suffix = true
	}

	if *passwordType == "string" {
		password, err := newRandomString(defaultOptions)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println(password[0])
	} else {
		fmt.Fprintf(os.Stderr, "ERR: Unsupported password type '%s'\n", *passwordType)
		os.Exit(1)
	}
}
