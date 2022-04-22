package main

import "testing"

func Test_validate(t *testing.T) {
	if err := validate(); err != nil {
		t.Errorf("Error while running test %v", err)
	}
}
