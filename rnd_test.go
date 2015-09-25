package rnd

import (
	"fmt"
	"hash"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"text/tabwriter"

	"github.com/jasonmoo/oc"
)

var (
	hashes = map[string]func() hash.Hash{
		"MD4":        MD4,
		"MD5":        MD5,
		"SHA1":       SHA1,
		"SHA224":     SHA224,
		"SHA256":     SHA256,
		"SHA384":     SHA384,
		"SHA512":     SHA512,
		"RIPEMD160":  RIPEMD160,
		"SHA3_224":   SHA3_224,
		"SHA3_256":   SHA3_256,
		"SHA3_384":   SHA3_384,
		"SHA3_512":   SHA3_512,
		"SHA512_224": SHA512_224,
		"SHA512_256": SHA512_256,
		"FNV1_64":    FNV1_64,
		"FNV1_64a":   FNV1_64a,
	}
	names = []string{
		"MD4",
		"MD5",
		"SHA1",
		"SHA224",
		"SHA256",
		"SHA384",
		"SHA512",
		"RIPEMD160",
		"SHA3_224",
		"SHA3_256",
		"SHA3_384",
		"SHA3_512",
		"SHA512_224",
		"SHA512_256",
		"FNV1_64",
		"FNV1_64a",
	}
)

func TestSource(t *testing.T) {

	const (
		set_size = 128
		runs     = 1 << 20
	)

	tabw := tabwriter.NewWriter(os.Stdout, 16, 8, 1, '\t', 0)
	defer tabw.Flush()

	fmt.Fprintln(tabw, "name\tmin\tmax\tdev\tdist\tmean\tstddev")

	seed := NewCryptoRandSeed()

	for _, name := range names {

		source := rand.New(NewHashSource(hashes[name], seed))

		set := oc.NewOc()
		var numbers []float64

		for i := 0; i < runs; i++ {
			numbers = append(numbers, float64(source.Int63()))
			set.Increment(strconv.Itoa(source.Intn(set_size)), 1)
		}

		if set.Len() != set_size {
			t.Errorf("Expected full distribution across set for %s, got %d", name, set.Len())
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

		stats := getStats(numbers)
		fmt.Fprintf(tabw, "%s\t%d\t%d\t%d\t%f\t%.0f\t%.0f\n", name, min, max, max-min, 1-(float64(max-min)/runs), stats.mean, stats.stdDev)

		tabw.Flush()
	}

	source := rand.New(rand.NewSource(seed))

	set := oc.NewOc()
	var numbers []float64

	for i := 0; i < runs; i++ {
		numbers = append(numbers, float64(source.Int63()))
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

	stats := getStats(numbers)
	fmt.Fprintf(tabw, "%s\t%d\t%d\t%d\t%f\t%.0f\t%.0f\n", "math/rand", min, max, max-min, 1-(float64(max-min)/runs), stats.mean, stats.stdDev)

}

func TestBench(_ *testing.T) {

	seed := NewCryptoRandSeed()

	for _, name := range names {

		source := rand.New(NewHashSource(hashes[name], seed))

		r := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				source.Int63()
			}
		})

		fmt.Println(name, r.String(), r.MemString())

	}

	source := rand.New(rand.NewSource(seed))

	r := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			source.Int63()
		}
	})

	fmt.Println("math/rand", r.String(), r.MemString())

}
