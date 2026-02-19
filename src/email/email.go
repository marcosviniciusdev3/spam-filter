package email

import (
	"errors"
	"io/fs"
	"strings"
)

type Email struct {
	FileName string
	Subject  string
	Body     string
}

const (
	ErrorFormat = "bad format on source file: "
)

func ReadEmailsFromFS(fileSystem fs.FS) ([]Email, error) {
	entries, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var emails []Email

	for _, v := range entries {
		data, err := fs.ReadFile(fileSystem, v.Name())
		if err != nil {
			return nil, err
		}
		e := Email{}
		e.FileName = v.Name()
		email, err := e.Parse(data)

		if err != nil {
			return nil, err
		}

		emails = append(emails, *email)
	}

	return emails, nil
}

// From from raw text
func (email *Email) Parse(src []byte) (*Email, error) {
	if string(src[:8]) != "Subject:" {
		return nil, errors.New(ErrorFormat)
	}

	var section_count = 0
	for i := range src {
		if rune(src[i]) == '\n' {
			section_count++
			// Parse Subject
			email.Subject = strings.TrimSpace(string(src[8:i]))
			// Parse Body
			if section_count == 1 {
				email.Body = strings.TrimSpace(string(src[i:]))
				break
			}
		}
	}

	return email, nil
}
