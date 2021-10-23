package worker

import ()

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}

type DatabaseHandler interface {
	Query(string, ...interface{}) (Row, error)
}