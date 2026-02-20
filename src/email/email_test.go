package email

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestParse_GivenRawTextEmail_ShouldReturnEmailVar(t *testing.T) {

	var stub_email = `Subject: my subject 
	lorem ipsum qui facit
	`
	var email Email

	got, err := email.Parse([]byte(stub_email))
	UnexpectedError(t, err)

	want := Email{
		FileName: "",
		Subject:  "my subject",
		Body:     "lorem ipsum qui facit",
	}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("got \"%+v\", want \"%+v\"", *got, want)
	}
}

// Test would fail if it encounters a directory
func TestReadEmailsFromFS_GivenOnlyFilesDataset_ShouldReturnEmailSlice(t *testing.T) {
	want := []Email{
		{FileName: "ham-1.md", Subject: "subject of ham-1.md", Body: "body 1"},
		{FileName: "spam-1.md", Subject: "subject of spam-1.md", Body: "body 2"},
	}

	fs := fstest.MapFS{
		want[0].FileName: {Data: []byte("Subject: " + want[0].Subject + "\n" + want[0].Body)},
		want[1].FileName: {Data: []byte("Subject: " + want[1].Subject + "\n" + want[1].Body)},
	}

	got, err := ReadEmailsFromFS(fs)
	UnexpectedError(t, err)

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

func TestNewEmailsFromFSError_WhenFSOpenFuncFails_ShouldFail(t *testing.T) {
	_, err := ReadEmailsFromFS(StubFailingFS{})

	if err == nil {
		t.Errorf("%s want an error, got nil", err)
	}
}

func TestDataSetTokenProbabity_GivenADataSet_ShouldReturnATokenProbability(t *testing.T) {
	dataset := []Email{
		{FileName: "ham-1.md", Subject: "subject of ham-1.md", Body: "body 1"},
		{FileName: "spam-1.md", Subject: "subject of spam-1.md", Body: "body 2"},
	}

	got, err := DataSetTokenProbability(dataset, []byte("body"))
	UnexpectedError(t, err)

	want := float32(0.2)

	if got != want {
		t.Errorf("got %#v, want %#v", got, want)
	}

}

func UnexpectedError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Unexpected error, got %q", err.Error())
	}
}
