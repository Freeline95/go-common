package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Transaction struct {
	*sqlx.Tx
}

func (t *Transaction) GetObject(qb sq.SelectBuilder, destObject interface{}) error {
	return GetObject(t, qb, destObject)
}
