package main

import "testing"

func TestCountWords(t *testing.T) {
	wds := words{[]string{"word", "word", "word", "wat"}}
	result := countWords(wds)

	if result.counts["word"] != 3 || result.counts["wat"] != 1 {
		t.Errorf("unexpected counts \"word\": %d \"wat\": %d",
			result.counts["word"], result.counts["wat"])
	}
}
