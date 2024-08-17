package urlshortener

import "testing"

func TestShortenUrl(t *testing.T) {
	longUrl := "url.com"
	got := shortenUrl(longUrl)
	want := "url.com"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
