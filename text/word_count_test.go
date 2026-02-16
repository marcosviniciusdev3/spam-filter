package text

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

var text = `Subject: subject 1
lorem ipsum qui facit facit`

func TestNewEmailsFromFS(t *testing.T) {
	fs := fstest.MapFS{
		"ham-1.md":  {Data: []byte(text)},
		"spam-1.md": {Data: []byte("I am spam!")},
	}

	emails, _ := NewEmailsFromFS(fs)

	if len(emails) != len(fs) {
		t.Errorf("got %d files, want %d files", len(emails), len(fs))
	}

}

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("I always fail!")
}

func TestNewEmailsFromFSError(t *testing.T) {
	_, err := NewEmailsFromFS(StubFailingFS{})

	if err == nil {
		t.Errorf("%s want an error, got nil", err)
	}
}

func TestWordCount(t *testing.T) {
	var w Word
	got, _ := w.Count([]byte(text))

	var want = Counter{
		{Word: "Subject", Num: 1},
		{Word: "subject", Num: 1},
		{Word: "1", Num: 1},
		{Word: "lorem", Num: 1},
		{Word: "ipsum", Num: 1},
		{Word: "qui", Num: 1},
		{Word: "facit", Num: 2},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got \"%+v\", want \"%+v\"", got, want)
	}
}
