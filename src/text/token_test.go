package text

import (
	"slices"
	"testing"

	"github.com/marcosviniciusdev3/spam-filter/lib/tests"
)

var mock_text = `Subject: subject 1
lorem qui ipsum qui facit facit facit`

var mock_counter = Counter{
	{Value: "Subject:", Num: 1},
	{Value: "subject", Num: 1},
	{Value: "1", Num: 1},
	{Value: "lorem", Num: 1},
	{Value: "ipsum", Num: 1},
	{Value: "qui", Num: 2},
	{Value: "facit", Num: 3},
}

func TestWordCount(t *testing.T) {
	got, _ := Count(slices.Clone([]byte(mock_text)))

	var want = slices.Clone(mock_counter)
	s := tests.Slice[Token]{}
	s.CompareUnorderedSlice(t, want, got)
}

// 1 - 0.01
func TestWordProbability(t *testing.T) {

	got, _ := Probability(slices.Clone([]byte(mock_text)), []byte("facit"))
	var want float32 = 0.3

	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}
