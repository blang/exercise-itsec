package main

import (
	"bytes"
	"fmt"
)

type Encoder func(s string, pw string) string

type Decoder func(s string, pw string) string

type Breaker interface {
	Break(s string) map[string]string
}

func main() {
	fmt.Printf("%c %d", 'a', 'a')
	fmt.Printf("%c %d", 'z', 'z')
	fmt.Printf("%s", CaesarEnc("hallo das ist ein test", "b"))
}

func CaesarEnc(s string, pw string) string {
	var pwshift int = int('a') - int(pw[0])
	fmt.Printf("Shift: %d\n", pwshift)
	var b bytes.Buffer
	for _, c := range s {
		if 'a'-c < 25 {
			b.WriteRune(rune((int('a')-int(c)+pwshift)%26 + int('a')))
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}
