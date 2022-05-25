package gorwos

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

type Options struct {
	Uppercase bool
	Lowercase bool
	Digits    bool
	Symbols   bool
	Prefix    bool
	Suffix    bool
	Length    int
	Count     int
}

func GenerateString(opts Options) ([]string, error) {
	// options := mergeWithDefaults(&opts)
	passwords, err := newRandomString(&opts)
	if err != nil {
		return nil, err
	}

	return passwords, nil
}

// func mergeWithDefaults(opts *Options) *Options {
// 	defaultOptions := &Options{
// 		Uppercase: true,
// 		Lowercase: true,
// 		Digits:    true,
// 		Symbols:   true,
// 		Prefix:    false,
// 		Suffix:    false,
// 		Length:    32,
// 		Count:     1,
// 	}

// 	if opts.Uppercase != nil {
// 		defaultOptions.Uppercase = opts.Uppercase
// 	}

// 	if opts.Lowercase != nil {
// 		defaultOptions.Lowercase = opts.Lowercase
// 	}

// 	if opts.Digits != nil {
// 		defaultOptions.Digits = opts.Digits
// 	}

// 	if opts.Symbols != nil {
// 		defaultOptions.Symbols = opts.Symbols
// 	}

// 	if opts.Prefix != nil {
// 		defaultOptions.Prefix = opts.Prefix
// 	}

// 	if opts.Suffix != nil {
// 		defaultOptions.Suffix = opts.Suffix
// 	}

// 	if opts.Length != nil {
// 		defaultOptions.Length = opts.Length
// 	}

// 	if opts.Count != nil {
// 		defaultOptions.Count = opts.Count
// 	}

// 	return defaultOptions
// }

func shuffleString(str string) []byte {
	rand.Seed(time.Now().Unix())

	data := []byte(str)
	rand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})

	return data
}

func generateKeyspace(opt *Options) []byte {
	var keyspace []string

	if opt.Uppercase {
		keyspace = append(keyspace, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}

	if opt.Lowercase {
		keyspace = append(keyspace, "abcdefghijklmnopqrstuvwxyz")
	}

	if opt.Digits {
		keyspace = append(keyspace, "0123456789")
	}

	if opt.Symbols {
		keyspace = append(keyspace, "@!-*#%=")
	}

	keyspaceJoined := strings.Join(keyspace, "")

	return shuffleString(keyspaceJoined)
}

func newRandomString(opt *Options) ([]string, error) {
	if !opt.Uppercase && !opt.Lowercase && !opt.Digits && !opt.Symbols {
		return nil, errors.New("err: You must specify atleast of one: uppercase, lowercase, digits or symbols")
	}

	keyspace := generateKeyspace(opt)
	var passwords []string

	for i := 0; i < opt.Count; i++ {
		var password []byte

		for j := 0; j < opt.Length; j++ {
			password = append(password, keyspace[rand.Intn(len(keyspace)-1)])
		}

		// TODO: Add prefix and suffix

		passwords = append(passwords, string(password))
	}

	return passwords, nil
}
