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

	srcIdx := rand.Intn(len(randomBytes))
	var chunkSize int
	var maxSize int

	for {
		chunkSize = size
		maxSize = len(randomBytes) - srcIdx
		if chunkSize > maxSize {
			chunkSize = maxSize
		}
		dest.Write(randomBytes[srcIdx : srcIdx+chunkSize])
		srcIdx += chunkSize
		if srcIdx >= len(randomBytes) {
			srcIdx = 0
		}
		if dest.Len() >= size {
			return
		}
	}
}
