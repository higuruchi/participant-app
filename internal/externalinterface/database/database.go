package database

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/higuruchi/participant-app/internal/interfaceadapter/repository/worker"
)

type DatabaseHandler struct {
	Conn *sql.DB
}

type TableRow struct {
	Rows *sql.Rows
}

type SQLResult struct {
	Result sql.Result
}

func NewDBHandler() (*DatabaseHandler, func()) {
	conn, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@(%s:%d)/%s",
		"user",
		"password",
		"192.168.0.104",
		3306,
		"participant-app",
	))
	if err != nil {
		log.Fatal("database connection error: ", err)
	}
	
	err = conn.Ping()
	if err!= nil {
		log.Fatal("database connection error %w", err)
	}
	return &DatabaseHandler{
		Conn: conn,
	}, func() { conn.Close() }
}

func (handler *DatabaseHandler) Query(
	statement string,
	args ...interface{},
) (worker.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return nil, fmt.Errorf("calling handler.Conn.Query: %w", err)
	}

	return &TableRow{
		Rows: rows,
	}, nil
}

func (tableRow *TableRow) Scan(dest ...interface{}) error {
	err := tableRow.Rows.Scan(dest...)
	if err != nil {
		return fmt.Errorf("calling handler.Rows.Scan: %w", err)
	}

	return nil
}

func (tableRow *TableRow) Next() bool {
	return tableRow.Rows.Next()
}

func (handler *DatabaseHandler) Execute(statement string, args ...interface{}) (worker.Result, error) {
	res := new(SQLResult)

	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return nil, err
	}

	res.Result = result
	return res, nil
}

func (r *SQLResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r *SQLResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

func (tableRow *TableRow) Close() error {
	err := tableRow.Rows.Close()
	if err != nil {
		return fmt.Errorf("calling tableRow.Row.Close: %w", err)
	}

	return nil
}