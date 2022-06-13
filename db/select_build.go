package db

import (
	"fmt"
	"strings"
)

func (r *dbSelect) buildSelect() string {
	return strings.Join(r.query, ", ")
}

func (r *dbSelect) buildFrom() string {
	return r.from
}

func (r *dbSelect) buildWhere() string {
	var where string
	return where
}

func (r *dbSelect) buildOrder() string {
	if len(r.order) == 0 {
		return ""
	}

	return fmt.Sprintf(" ORDER BY %s", strings.Join(r.order, ", "))
}

func (r *dbSelect) buildJoin() string {
	var join string

	if len(r.join) == 0 {
		return ""
	}

	for _, item := range r.join {
		join += fmt.Sprintf(" %s JOIN %s ON %s", item.joinType, item.join, item.on)
	}

	return join
}

func (r *dbSelect) buildLimit() string {
	var (
		limit  string
		offset string
	)

	if r.limit.offset != nil {
		offset = fmt.Sprintf(" OFFSET %d", *r.limit.offset)
	}

	if r.limit.limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", r.limit.limit)
	}

	return fmt.Sprintf("%s%s", offset, limit)
}
