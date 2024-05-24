package db

import (
	"context"
	"fmt"

	"database/sql"

	common_error "github.com/Freeline95/go-common/pkg/error"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type SQLDB interface {
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	DriverName() string
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	MustExec(query string, args ...interface{}) sql.Result
	MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	Preparex(query string) (*sqlx.Stmt, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	Rebind(query string) string
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...any) (sql.Result, error)
	GetObject(qb sq.SelectBuilder, destObject interface{}) error
}

type DB struct {
	*sqlx.DB
}

func NewDB(dbHost string, dbPort int, dbUser, dbPassword, dbName, dbDriverName string) (*DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)

	db, err := sqlx.Connect(dbDriverName, connStr)
	if err != nil {
		return nil, common_error.Annotate(err, "Error while connect to DB")
	}

	// Check connection with DB
	err = db.Ping()
	if err != nil {
		return nil, common_error.Annotate(err, "Error while ping DB")
	}

	log.Println("Connected to the database successfully")

	return &DB{db}, nil
}

func (db *DB) Use(tx *Transaction) SQLDB {
	if tx == nil {
		return db
	}

	return tx
}

func (db *DB) BeginTransaction() (*Transaction, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, common_error.Annotate(err, "Error while beginx")
	}

	return &Transaction{tx}, nil
}

func (db *DB) GetObject(qb sq.SelectBuilder, destObject interface{}) error {
	return GetObject(db, qb, destObject)
}

func GetObject(q sqlx.Queryer, qb sq.SelectBuilder, destObject interface{}) error {
	sql, args, err := qb.ToSql()
	if err != nil {
		return common_error.Annotate(err, "Error while to sql")
	}

	rows, err := q.Queryx(sql, args...)
	if err != nil {
		return common_error.Annotate(err, "Error queryx")
	}
	defer rows.Close()

	if err := rows.StructScan(&destObject); err != nil {
		return common_error.Annotate(err, "Error scanning object")
	}

	return nil
}
