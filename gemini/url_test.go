package gemini

import (
	"fmt"
	"testing"
)

var urls []string = []string{
	"gemini.circumlunar.space",
	"gemini.circumlunar.space:1965",
	"gemini://gemini.circumlunar.space",
	"gemini://gemini.circumlunar.space:1965",
	"docs/",
}

func testUrl(url string) {

}
func TestUrl(t *testing.T) {

	for i, url := range urls {
		t.Run(fmt.Sprint("Test URL %d", i), testUrl(url))
	}
}
