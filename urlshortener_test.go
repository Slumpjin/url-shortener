package urlshortener

import (
	"fmt"
	"reflect"
	"testing"
)

func TestShortenUrlPath(t *testing.T) {
	t.Run("test url path returning base 10 id", func(t *testing.T) {
		urlPath := "a"
		id, err := shortenUrlPath(urlPath)
		got := reflect.TypeOf(id)
		want := reflect.TypeOf(int64(0))

		if got != want && err == nil {
			t.Errorf("got %T want %T", id, int(0))
		}
	})
	t.Run("test base 36 id", func(t *testing.T) {
		urlPath := "rs"
		got, err := shortenUrlPath(urlPath)
		want := int64(1000)
		fmt.Println(got, want)

		if got != want && err == nil {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("test slightly longer base 36 id", func(t *testing.T) {
		urlPath := "lfls"
		got, err := shortenUrlPath(urlPath)
		want := int64(1000000)

		if got != want && err == nil {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
