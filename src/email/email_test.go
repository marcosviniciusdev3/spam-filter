package email

import (
	"errors"
	"fmt"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

var stub_email = `Subject: my subject 
lorem ipsum qui facit
`

func Test(t *testing.T) {
	var email Email

	got, _ := email.Parse([]byte(stub_email))

	want := Email{
		FileName: "",
		Subject:  "my subject",
		Body:     "lorem ipsum qui facit",
	}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("got \"%+v\", want \"%+v\"", *got, want)
	}
}

func TestReadEmailsFromFS(t *testing.T) {
	want := []Email{
		{FileName: "ham-1.md", Subject: "subject of ham-1.md", Body: "body 1"},
		{FileName: "spam-1.md", Subject: "subject of spam-1.md", Body: "body 2"},
	}

	fs := fstest.MapFS{
		want[0].FileName: {Data: []byte("Subject: " + want[0].Subject + "\n" + want[0].Body)},
		want[1].FileName: {Data: []byte("Subject: " + want[1].Subject + "\n" + want[1].Body)},
	}

	got, err := ReadEmailsFromFS(fs)

	if err != nil {
		fmt.Printf("%v", err)
	}

	if len(got) != len(want) {
		t.Errorf("got %d files, want %d files", len(got), len(want))
	}

	// The order is important here
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}

}

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("I always fail!")
}

func TestNewEmailsFromFSError(t *testing.T) {
	_, err := ReadEmailsFromFS(StubFailingFS{})

	if err == nil {
		t.Errorf("%s want an error, got nil", err)
	}
}
