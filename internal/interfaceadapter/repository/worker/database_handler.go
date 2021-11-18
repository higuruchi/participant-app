package worker

import ()

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}

type Result interface {
	LastInsertId()(int64, error)
	RowsAffected() (int64, error)
}

type DatabaseHandler interface {
	Query(string, ...interface{}) (Row, error)
	Execute(statement string, args ...interface{}) (Result, error)
}