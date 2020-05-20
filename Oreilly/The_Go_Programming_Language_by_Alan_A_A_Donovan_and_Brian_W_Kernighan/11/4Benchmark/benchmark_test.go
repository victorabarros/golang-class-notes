package benchmark

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkMapArrayToString(b *testing.B) {
	attempt := struct {
		arr []int
		sep string
	}{
		arr: []int{1, 0, 5},
		sep: ".",
	}
	for ii := 0; ii < b.N; ii++ {
		mapArrayToString(attempt.arr, attempt.sep)
	}
}

func BenchmarkNMillisecondsSleep(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		nMillisecondsSleep(10)
	}
}

func BenchmarkNNanosecondsSleep(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		nNanosecondsSleep(10)
	}
}

func BenchmarkIsPalindromeV1(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		isPalindromeV1("A man, a plan, a canal: Panama")
	}
}

func BenchmarkIsPalindromeV2(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		isPalindromeV2("A man, a plan, a canal: Panama")
	}
}

func buildRandomPalindrome(length int) string {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	runes := make([]rune, length)

    for i := 0; i < (length+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[length-1-i] = r
	}

	return string(runes)
}

func benchmarkIsPalindromeV1RandomWord(b *testing.B, length int) {
	randomPalindrome := buildRandomPalindrome(length)

	for ii := 0; ii < b.N; ii++ {
		isPalindromeV2(randomPalindrome)
	}
}

func BenchmarkIsPalindromeV1RandomWord10(b *testing.B)   { benchmarkIsPalindromeV1RandomWord(b, 10) }
func BenchmarkIsPalindromeV1RandomWord100(b *testing.B)  { benchmarkIsPalindromeV1RandomWord(b, 100) }
func BenchmarkIsPalindromeV1RandomWord1000(b *testing.B) { benchmarkIsPalindromeV1RandomWord(b, 1000) }

func TestIsPalindromeV1(t *testing.T) {
	attemps := []struct {
		word string
		want bool
	}{
		{
			word: "A man, a plan, a canal: Panama",
			want: false,
		},
		{
			word: "tieit",
			want: true,
		},
		{
			word: "blaalb",
			want: true,
		},
		{
			word: "hotel-worker",
			want: false,
		},
	}

	for _, attemp := range attemps {
		if isPalindromeV1(attemp.word) != attemp.want {
			t.Error(attemp)
		}
	}
}

func TestIsPalindromeV2(t *testing.T) {
	attemps := []struct {
		word string
		want bool
	}{
		{
			word: "A man, a plan, a canal: Panama",
			want: false,
		},
		{
			word: "tieit",
			want: true,
		},
		{
			word: "blaalb",
			want: true,
		},
		{
			word: "hotel-worker",
			want: false,
		},
	}

	for _, attemp := range attemps {
		if isPalindromeV2(attemp.word) != attemp.want {
			t.Error(attemp)
		}
	}
}
