package rnd

import (
	"crypto"
	"fmt"
	mrand "math/rand"
	"os"
	"strconv"
	"testing"
	"text/tabwriter"

	"github.com/jasonmoo/oc"
)

func TestInt(t *testing.T) {

	const (
		set_size = 10
		runs     = 1 << 20
	)

	hashes := map[string]crypto.Hash{
		"MD4":        crypto.MD4,        // import golang.org/x/crypto/md4
		"MD5":        crypto.MD5,        // import crypto/md5
		"SHA1":       crypto.SHA1,       // import crypto/sha1
		"SHA224":     crypto.SHA224,     // import crypto/sha256
		"SHA256":     crypto.SHA256,     // import crypto/sha256
		"SHA384":     crypto.SHA384,     // import crypto/sha512
		"SHA512":     crypto.SHA512,     // import crypto/sha512
		"RIPEMD160":  crypto.RIPEMD160,  // import golang.org/x/crypto/ripemd160
		"SHA3_224":   crypto.SHA3_224,   // import golang.org/x/crypto/sha3
		"SHA3_256":   crypto.SHA3_256,   // import golang.org/x/crypto/sha3
		"SHA3_384":   crypto.SHA3_384,   // import golang.org/x/crypto/sha3
		"SHA3_512":   crypto.SHA3_512,   // import golang.org/x/crypto/sha3
		"SHA512_224": crypto.SHA512_224, // import crypto/sha512
		"SHA512_256": crypto.SHA512_256, // import crypto/sha512
	}

	tabw := tabwriter.NewWriter(os.Stdout, 16, 8, 1, '\t', 0)
	defer tabw.Flush()

	fmt.Fprintln(tabw, "name\tmin\tmax\tdev\tdist")

	for name, h := range hashes {

		rand := NewSource(h)
		// rand.Seed([]byte("jason"))

		set := oc.NewOc()

		for i := 0; i < runs; i++ {
			set.Increment(strconv.Itoa(rand.Intn(set_size)), 1)
		}

		set.SortByCt(oc.DESC)

		var min, max int64
		for set.Next() {
			_, v := set.KeyValue()
			if min == 0 || v < min {
				min = v
			}
			if v > max {
				max = v
			}
			// fmt.Println(k, "\t", v)
		}

		fmt.Fprintf(tabw, "%s\t%d\t%d\t%d\t%f\n", name, min, max, max-min, 1-(float64(max-min)/runs))
		tabw.Flush()
	}

	source := mrand.New(mrand.NewSource(NewSource(crypto.SHA3_512).Int63()))

	set := oc.NewOc()

	for i := 0; i < runs; i++ {
		set.Increment(strconv.Itoa(source.Intn(set_size)), 1)
	}

	set.SortByCt(oc.DESC)

	var min, max int64
	for set.Next() {
		_, v := set.KeyValue()
		if min == 0 || v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	fmt.Fprintf(tabw, "%s\t%d\t%d\t%d\t%f\n", "math/rand", min, max, max-min, 1-(float64(max-min)/runs))

}

func BenchmarkMD4(b *testing.B) {
	h := NewSource(crypto.MD4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Int63()
	}
}
func BenchmarkMD5(b *testing.B) {
	h := NewSource(crypto.MD5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Int63()
	}
}
func BenchmarkSHA1(b *testing.B) {
	h := NewSource(crypto.SHA1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Int63()
	}
}
func BenchmarkSHA224(b *testing.B) {
	h := NewSource(crypto.SHA224)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Int63()
	}
}
func BenchmarkSHA256(b *testing.B) {
	h := NewSource(crypto.SHA256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Int63()
	}
}
func BenchmarkSHA384(b *testing.B) {
	h := NewSource(crypto.SHA384)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Int63()
	}
}
func BenchmarkSHA512(b *testing.B) {
	h := NewSource(crypto.SHA512)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Int63()
	}
}
func BenchmarkRIPEMD160(b *testing.B) {
	h := NewSource(crypto.RIPEMD160)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Int63()
	}
}
func BenchmarkSHA3_224(b *testing.B) {
	h := NewSource(crypto.SHA3_224)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Int63()
	}
}
func BenchmarkSHA3_256(b *testing.B) {
	h := NewSource(crypto.SHA3_256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Int63()
	}
}
func BenchmarkSHA3_384(b *testing.B) {
	h := NewSource(crypto.SHA3_384)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Int63()
	}
}
func BenchmarkSHA3_512(b *testing.B) {
	h := NewSource(crypto.SHA3_512)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Int63()
	}
}
func BenchmarkSHA512_224(b *testing.B) {
	h := NewSource(crypto.SHA512_224)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Int63()
	}
}
func BenchmarkSHA512_256(b *testing.B) {
	h := NewSource(crypto.SHA512_256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Int63()
	}
}
func BenchmarkMathRand(b *testing.B) {
	source := mrand.New(mrand.NewSource(NewSource(crypto.SHA3_512).Int63()))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		source.Int63()
	}
}