package crud

import (
	"fmt"
	"testing"
)

type fakeDB struct{}

func (fk fakeDB) Create(u *User) (string, error) {
	return fmt.Sprintf("Created: %v", u), nil
}

func testCreate(t *testing.T) {
	q := query{
		inm: fakeDB{},
	}
	q.Query("Create admin admin@example.com")
	msg, err := q.Create(&User{"admin", "admin@example.com"})
	exp := "Created: {admin admin@example.com}"
	if msg != exp {
		t.Fatalf("Expected: %q, got: %q", exp, msg)
	}
}
