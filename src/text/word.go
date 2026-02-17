package text

import (
	"io/fs"

	"github.com/marcosviniciusdev3/spam-filter/email"
)

type Counter []Word

var counter Counter

type Word struct {
	Word string
	Num  uint
}

func (w *Word) Count(text []byte) (Counter, error) {

	return counter, nil
}

func NewEmailsFromFS(fileSystem fs.FS) ([]email.Email, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var emails []email.Email
	for range dir {
		emails = append(emails, email.Email{})
	}
	return emails, nil
}
