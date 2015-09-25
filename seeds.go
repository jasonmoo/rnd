package rnd

import (
	"crypto/rand"
	"unsafe"
)

func NewCryptoRandBytes(size int) []byte {
	seed := make([]byte, size)
	rand.Read(seed)
	return seed
}

func NewCryptoRandSeed() int64 {
	seed := make([]byte, 8)
	rand.Read(seed)
	return *(*int64)(unsafe.Pointer(&seed[0]))
}
