package postgres

import (
	"github.com/uptrace/bun"
	"oss-backend/internal/persistence"
)

func (p *Postgres) apply(opts []persistence.QueryBuilder) func(bun.QueryBuilder) bun.QueryBuilder {
	return func(idb bun.QueryBuilder) bun.QueryBuilder {
		for i := range opts {
			idb = opts[i](idb)
		}

		return idb
	}
}

func (p *Postgres) selectApply(opts []persistence.SelectOption) func(*bun.SelectQuery) *bun.SelectQuery {
	return func(query *bun.SelectQuery) *bun.SelectQuery {
		for i := range opts {
			query.Apply(opts[i])
		}

		return query
	}
}
