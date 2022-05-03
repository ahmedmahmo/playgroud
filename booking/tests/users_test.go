package tests

import (
	"testing"
	"github.com/ahmedmahmo/learn/booking/db"
)

func TestChecksValidation(t *testing.T) {
	u := &db.User{}
	err := u.Validate()
	if err != nil {
		t.Fatal(err)
	}
}