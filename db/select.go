package db

import (
	"fmt"
	"strings"
)

type dbSelect struct {
	query []string
	order []string
	where []whereDb
	from  string
}

func NewSelect(queryString string) *dbSelect {
	return &dbSelect{
		query: strings.Split(queryString, ","),
	}
}

func (r *dbSelect) From(from string) *dbSelect {
	r.from = from
	return r
}

func (r *dbSelect) OrderBy(order string) *dbSelect {
	r.order = append(r.order, strings.Split(order, ",")...)
	return r
}

func (r *dbSelect) Where(column string, value interface{}) *dbSelect {
	r.where = append(r.where, whereDb{
		column:    column,
		value:     value,
		separator: "=",
	})
	return r
}

func (r *dbSelect) WhereSeparator(column string, separator string, value interface{}) *dbSelect {
	r.where = append(r.where, whereDb{
		column:    column,
		value:     value,
		separator: separator,
	})
	return r
}

func (r *dbSelect) Select(selectQuery string) *dbSelect {
	r.query = append(r.query, strings.Split(selectQuery, ",")...)
	return r
}

func (r *dbSelect) Get() string {
	return fmt.Sprintf(
		"SELECT %s FROM %s%s%s%s%s",
		r.buildSelect(),
		r.buildFrom(),
		r.buildJoin(),
		r.buildWhere(),
		r.buildOrder(),
		r.buildLimit(),
	)
}
