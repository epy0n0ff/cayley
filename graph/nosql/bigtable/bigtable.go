package bigtable

import (
	"context"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/graph/nosql"
)

const Type = "mongo"

var (
	_ nosql.Database = (*DB)(nil)
)

type DB struct {
}

func (db *DB) Insert(ctx context.Context, col string, key nosql.Key, d nosql.Document) (nosql.Key, error) {
	panic("implement me")
}

func (db *DB) FindByKey(ctx context.Context, col string, key nosql.Key) (nosql.Document, error) {
	panic("implement me")
}

func (db *DB) Query(col string) nosql.Query {
	panic("implement me")
}

func (db *DB) Update(col string, key nosql.Key) nosql.Update {
	panic("implement me")
}

func (db *DB) Delete(col string) nosql.Delete {
	panic("implement me")
}

func (db *DB) EnsureIndex(ctx context.Context, col string, primary nosql.Index, secondary []nosql.Index) error {
	panic("implement me")
}

func (db *DB) Close() error {
	panic("implement me")
}

func init() {
	nosql.Register(Type, nosql.Registration{
		NewFunc:      Open,
		InitFunc:     Create,
		IsPersistent: true,
	})
}

func Open(string, graph.Options) (nosql.Database, error) {

}

func Create(string, graph.Options) (nosql.Database, error) {

}
