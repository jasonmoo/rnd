package rnd

import (
	"crypto"
	"hash"
	"hash/fnv"
	"unsafe"

	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"

	_ "golang.org/x/crypto/md4"
	_ "golang.org/x/crypto/ripemd160"
	_ "golang.org/x/crypto/sha3"
)

var (
	MD4         = crypto.MD4.New
	MD5         = crypto.MD5.New
	SHA1        = crypto.SHA1.New
	SHA224      = crypto.SHA224.New
	SHA256      = crypto.SHA256.New
	SHA384      = crypto.SHA384.New
	SHA512      = crypto.SHA512.New
	RIPEMD160   = crypto.RIPEMD160.New
	SHA3_224    = crypto.SHA3_224.New
	SHA3_256    = crypto.SHA3_256.New
	SHA3_384    = crypto.SHA3_384.New
	SHA3_512    = crypto.SHA3_512.New
	SHA512_224  = crypto.SHA512_224.New
	SHA512_256  = crypto.SHA512_256.New
	FNV1_64     = func() hash.Hash { return fnv.New64().(hash.Hash) }
	FNV1_64a    = func() hash.Hash { return fnv.New64a().(hash.Hash) }
	MyShiftHash = func() hash.Hash { return new(MyShift) }
)

type (
	MyShift int64
)

func (m *MyShift) Write(b []byte) (int, error) {
	b_len := len(b)
	for len(b) > 8 {
		n := *(*int64)(unsafe.Pointer(&b[0]))
		em := int64(*m)
		*m = MyShift((em * n) ^ (n << 8) ^ (em >> 6))
		b = b[8:]
	}
	buf := [8]byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA}
	copy(buf[:], b)
	n := *(*int64)(unsafe.Pointer(&buf[0]))
	em := int64(*m)
	*m = MyShift((em * n) ^ (n << 8) ^ (em >> 6))
	return b_len, nil
}

func (m *MyShift) Sum(b []byte) []byte {
	out := *(*[8]byte)(unsafe.Pointer(m))
	return append(b, out[0], out[1], out[2], out[3], out[4], out[5], out[6], out[7])
}

func (m *MyShift) Size() int      { return 8 }
func (m *MyShift) BlockSize() int { return 8 }
func (m *MyShift) Reset()         {}
