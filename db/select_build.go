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
	return ""
}

func (r *dbSelect) buildOrder() string {
	if len(r.order) == 0 {
		return ""
	}

	return fmt.Sprintf(" ORDER BY %s", strings.Join(r.order, ", "))
}

func (r *dbSelect) buildJoin() string {
	return ""
}

func (r *dbSelect) buildLimit() string {
	return ""
}
