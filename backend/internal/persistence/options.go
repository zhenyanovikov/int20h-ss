package persistence

import (
	"strings"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type QueryBuilder func(idb bun.QueryBuilder) bun.QueryBuilder
type SelectOption func(query *bun.SelectQuery) *bun.SelectQuery

func Like(fieldName string, title string) QueryBuilder {
	return func(query bun.QueryBuilder) bun.QueryBuilder {
		text := "%" + strings.ToLower(title) + "%"
		return query.Where("lower("+fieldName+") ILIKE ?", text)
	}
}

func Exact(fieldName string, data any) QueryBuilder {
	return func(query bun.QueryBuilder) bun.QueryBuilder {
		return query.Where(fieldName+" = ?", data)
	}
}

func WithIDIn(ids []uuid.UUID) QueryBuilder {
	return func(query bun.QueryBuilder) bun.QueryBuilder {
		return query.Where("student.id IN (?)", bun.In(ids))
	}
}
