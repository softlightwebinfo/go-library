package test

import (
	"fmt"
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

func TestSelectJoinDb(t *testing.T) {
	orm := db.NEW()

	query := orm.Select("u.id,u.name").
		From("users u").
		Select("u.email").
		Join("categories c", "c.id = u.fk_category_id").
		JoinType("roles r", "RIGHT", "r.id=u.fk_role_id").
		Where("u.id", 12).
		OrderBy("u.id desc").
		LimitOffset(0, 10).
		Get()

	fmt.Println(query)
}
