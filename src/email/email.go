package email

import (
	"errors"
	"strings"
)

type Email struct {
	Subject string
	Body    string
}

const (
	ErrorFormat = "bad format on source file: "
)

// From from raw text
func (email *Email) Parse(src []byte) (*Email, error) {
	if string(src[:8]) != "Subject:" {
		return &Email{}, errors.New(ErrorFormat)
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

	println(section_count)

	return email, nil
}
