package random

import (
	"bytes"
	"math/rand"
)

//goland:noinspection SpellCheckingInspection
const baseCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var randomBytes []byte

func init() {
	Randomize(128)
}

func Randomize(size int) {
	if size < 1 {
		size = 1
	}
	randomBytes = make([]byte, size)
	for i := range randomBytes {
		randomBytes[i] = baseCharset[rand.Intn(len(baseCharset))]
	}
}

func Write(dest *bytes.Buffer, size int) {
	if size < 1 {
		return
	}

	// Randomize initial position to copy bytes from.
	pos := rand.Intn(len(randomBytes))

	var chunkSize int

	for {
		chunkSize = size
		if chunkSize > len(randomBytes)-pos {
			chunkSize = len(randomBytes) - pos
		}
		dest.Write(randomBytes[pos : pos+chunkSize])
		size -= chunkSize
		if size <= 0 {
			return
		}
		pos += chunkSize
		if pos >= len(randomBytes) {
			pos = 0
		}
	}
}
