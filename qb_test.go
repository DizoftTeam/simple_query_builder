package sqb

import (
	"testing"
)

func TestSimpleGenerate(t *testing.T) {
	query := NewQuery("users as u").Generate()
	must := "SELECT * FROM users as u WHERE 1 = 1"

	if query != must {
		t.Errorf("Query generated uncorrectly!\nWant '%v', but get '%v'", must, query)
	}
}

func TestWhere(t *testing.T) {
	query := NewQuery("users as u").
		Where("and", "u.name = 'User'").
		Where("and", "u.lastname = 'User1'").
		Generate()

	must := "SELECT * FROM users as u WHERE 1 = 1 AND u.name = 'User' AND u.lastname = 'User1'"

	if query != must {
		t.Errorf("Query generated uncorrectly!\nWant '%v', but get '%v'", must, query)
	}
}
