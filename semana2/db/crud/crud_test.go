package crud

import (
	"testing"
)

const (
	createm = "Created: &{admin admin@example.com}\nIn DB: map[1:{admin admin@example.com}]"
	retm    = "Value from map: &{admin admin@example.com}"
	updm    = "Updated, old:&{admin admin@example.com}, new:&{admin1 admin1@example.com}\nNew map: map[1:{admin1 admin1@example.com}]"
	delm    = "Deleted: &{admin1 admin1@example.com}\nIn DB: map[]"
)

var q = Query{
	Inm: Database{
		DB: make(map[int]*User),
	},
}

func TestCrud(t *testing.T) {

	u := User{"admin", "admin@example.com"}
	msg, _ := q.Inm.Create(&u)
	if msg != createm {
		t.Errorf("Wanted: %q, Got %q", msg, createm)
	}
	r, _ := q.Inm.Retrieve(1)
	if r != retm {
		t.Errorf("Wanted: %q, Got %q", retm, r)
	}
	up, _ := q.Inm.Update(&User{"admin1", "admin1@example.com"}, 1)
	if up != updm {
		t.Errorf("Wanted: %q, Got %q", updm, up)
	}
	d, _ := q.Inm.Del(1)
	if d != delm {
		t.Errorf("Wanted: %q, Got %q", delm, d)
	}
}

func TestCrudErrors(t *testing.T) {
	u := User{"user", "user@example.com"}
	msg, _ := q.Inm.Create(&u)
	exp := "Created: &{user user@example.com}\nIn DB: map[2:{user user@example.com}]"
	if msg != exp {
		t.Errorf("Wanted: %q, Got %q", exp, msg)
	}
	_, e := q.Inm.Retrieve(3)
	if e == nil {
		t.Error(e)
	}
	_, e1 := q.Inm.Update(&User{"admin1", "admin1@example.com"}, 3)
	if e1 == nil {
		t.Error(e1)
	}
	_, e2 := q.Inm.Del(3)
	if e2 == nil {
		t.Error(e2)
	}
}
