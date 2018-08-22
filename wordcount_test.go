package main

import "testing"

type testCase struct {
	name  string
	words words
}

var cases = []testCase{
	{name: "4 words", words: words{[]string{"word", "word", "word", "wat"}}},
	{name: "8 words", words: words{[]string{"word", "word", "word", "wat", "word", "word", "word", "wat"}}},
	{name: "16 words", words: words{[]string{"word", "word", "word", "wat", "word", "word", "word", "wat", "word", "word", "word", "wat"}}},
}

func TestCountWords(t *testing.T) {
	wds := words{[]string{"word", "word", "word", "wat"}}
	result := countWords(wds)

	if result.Counts["word"] != 3 || result.Counts["wat"] != 1 {
		t.Errorf("unexpected counts \"word\": %d \"wat\": %d",
			result.Counts["word"], result.Counts["wat"])
	}
}

func BenchmarkCountWords(b *testing.B) {
	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Sometimes the compiler will think it can optimize calls
				// it thinks are no-ops entirely out of the asm, but this one
				// doesn't seem to so it's fine
				countWords(c.words)
			}
		})
	}
}
