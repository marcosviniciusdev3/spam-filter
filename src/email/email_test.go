package email

import (
	"reflect"
	"testing"
)

var stub_email = `Subject: my subject 
lorem ipsum qui facit
`

func Test(t *testing.T) {
	var email Email

	got, _ := email.Parse([]byte(stub_email))

	want := Email{
		Subject: "my subject",
		Body:    "lorem ipsum qui facit",
	}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("got \"%+v\", want \"%+v\"", *got, want)
	}
}
