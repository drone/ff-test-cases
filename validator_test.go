package validator

import "testing"

func Test_validate(t *testing.T) {
	err := validate()
	if err != nil {
		t.Fatal(err)
	}
}
