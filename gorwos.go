package main

import (
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

func newRandomString(opt *options) []string {
	keyspace := generateKeyspace(opt)
	var passwords []string
	var password []byte

	for i := 0; i < opt.length; i++ {
		password = append(password, keyspace[rand.Intn(len(keyspace)-1)])
	}

	passwords = append(passwords, string(password))

	return passwords
}

func main() {

	var length = flag.Int("l", 16, "length of string")
	var passwordType = flag.String("t", "string", "type of password to generate")
	flag.Parse()

	defaultOptions := &options{
		uppercase: true,
		lowercase: true,
		digits:    true,
		symbols:   true,
		length:    *length,
		count:     1,
	}

	if *passwordType == "string" {
		fmt.Println(newRandomString(defaultOptions)[0])
	} else {
		fmt.Fprintf(os.Stderr, "Unsupported password type '%s'\n", *passwordType)
		os.Exit(1)
	}
}
