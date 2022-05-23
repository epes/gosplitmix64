package gosplitmix64_test

import (
	"github.com/epes/gosplitmix64"
	"math"
	"testing"
)

func TestNext(t *testing.T) {
	r := gosplitmix64.New(1234567)

	tests := []uint64{
		6457827717110365317,
		3203168211198807973,
		9817491932198370423,
		4593380528125082431,
		16408922859458223821,
	}

	for _, expected := range tests {
		got := r.Next()

		if expected != got {
			t.Fatalf("expected: %v, got: %v", expected, got)
		}
	}
}

func TestSpread(t *testing.T) {
	r := gosplitmix64.New(987654321)
	results := make([]uint64, 5)

	for i := 0; i < 100000; i++ {
		asRatio := float64(r.Next()) / (1 << 64)
		floor := int(math.Floor(asRatio * 5))
		results[floor]++
	}

	tests := []uint64{
		20027,
		19892,
		20073,
		19978,
		20030,
	}

	for idx, expected := range tests {
		got := results[idx]

		if expected != got {
			t.Fatalf("expected: %v, got: %v", expected, got)
		}
	}
}
