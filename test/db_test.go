package test

import (
	"testing"

	"github.com/softlightwebinfo/go-library/db"
)

func TestSelectDb(t *testing.T) {
	orm := db.NEW()

	query := orm.Select("u.id,u.name").
		From("users u").
		Select("u.email").
		Where("u.id", 12).
		OrderBy("u.id desc").
		Get()

	if query != "SELECT u.id, u.name, u.email FROM users u ORDER BY u.id desc" {
		t.Fatalf("error: (%s) QUERY MAPPER [%s]", "SELECT u.id, u.name, u.email FROM users u ORDER BY u.id desc", query)
	}
}
