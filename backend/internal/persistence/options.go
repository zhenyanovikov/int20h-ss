package persistence

import (
	"strings"

	"github.com/uptrace/bun"
)

type QueryBuilder func(idb bun.QueryBuilder) bun.QueryBuilder
type SelectOption func(query *bun.SelectQuery) *bun.SelectQuery

func WithTitleLike(title string) QueryBuilder {
	return func(query bun.QueryBuilder) bun.QueryBuilder {
		text := "%" + strings.ToLower(title) + "%"
		return query.Where("lower(title) ILIKE ?", text)
	}
}
