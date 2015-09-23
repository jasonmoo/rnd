package rnd

import (
	"crypto"
	"crypto/rand"
	"hash"
	"unsafe"

	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"

	_ "golang.org/x/crypto/md4"
	_ "golang.org/x/crypto/ripemd160"
	_ "golang.org/x/crypto/sha3"
)

type (
	Source struct {
		h    hash.Hash
		seed []byte
	}
)

func NewSource(h crypto.Hash) *Source {
	seed := make([]byte, 8)
	rand.Read(seed)
	return &Source{
		seed: seed,
		h:    h.New(),
	}
}

func (s *Source) Seed(seed []byte) {
	s.h.Write(seed)
	s.seed = s.h.Sum(nil)
}

func (s *Source) Int() int {
	return int(s.Int63())
}

func (s *Source) Intn(n int) int {
	v := s.Int()
	if mask := n - 1; n&mask == 0 { // n is power of two, can mask
		return v & mask
	}
	// find evenly divisible max and look for
	// number within it
	const (
		max_val = (^uint(0) >> 2)
		bound   = ^(^uint(0) >> 1)
	)
	max := int(max_val - (bound % uint(n)))
	for v > max {
		v = s.Int()
	}
	return v % n
}

func (s *Source) Int63() int64 {
	s.hash()
	const mask = int64(^uint64(0) >> 1)
	return *(*int64)(unsafe.Pointer(&s.seed[0])) & mask
}

func (s *Source) hash() {
	s.h.Write(s.seed)
	copy(s.seed, s.h.Sum(nil)[:8])
}
