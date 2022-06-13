package db

import (
	"fmt"
	"strings"
)

const (
	SELECT_JOIN_LEFT  = "LEFT"
	SELECT_JOIN_RIGHT = "RIGHT"
	SELECT_JOIN       = ""
	SELECT_JOIN_INNER = "INNER"
)

type dbSelect struct {
	query []string
	from  string
	order []string
	where []whereDb
	join  []joinDb
	limit limitDb
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

func (r *dbSelect) Join(join string, on string) *dbSelect {
	r.JoinType(join, SELECT_JOIN, on)
	return r
}

func (r *dbSelect) JoinRight(join string, on string) *dbSelect {
	r.JoinType(join, SELECT_JOIN_RIGHT, on)
	return r
}

func (r *dbSelect) JoinInner(join string, on string) *dbSelect {
	r.JoinType(join, SELECT_JOIN_INNER, on)
	return r
}

func (r *dbSelect) JoinType(join string, joinType string, on string) *dbSelect {
	r.join = append(r.join, joinDb{
		join:     join,
		on:       on,
		joinType: joinType,
	})
	return r
}

func (r *dbSelect) Limit(limit int) *dbSelect {
	r.limit = limitDb{
		limit: limit,
	}
	return r
}

func (r *dbSelect) Offset(offset int) *dbSelect {
	r.limit = limitDb{
		offset: &offset,
	}
	return r
}

func (r *dbSelect) LimitOffset(offset int, limit int) *dbSelect {
	r.limit = limitDb{
		offset: &offset,
		limit:  limit,
	}
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
