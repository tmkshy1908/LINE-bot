package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// var db *sql.DB

type SqlHandler struct {
	Conn *sql.DB
}
type dbSettings struct {
	User     string
	Dbname   string
	Passwrod string
	Database string
}

type SqlConnHandler interface {
	Exec(context.Context, string, ...interface{}) (sql.Result, error)
	Query(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRow(context.Context, string, ...interface{}) *sql.Row
	ExecWithTx(txFunc func(*sql.Tx) error) error
}

func NewHandler() (h SqlConnHandler, err error) {
	conf := dbSettings{
		User:     "yamadatarou",
		Dbname:   "bot_schedule",
		Passwrod: "1234",
		// Database: "database",
	}
	connectionString := fmt.Sprintf(conf.User, conf.Dbname, conf.Passwrod)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
	}

	myq := &SqlHandler{Conn: db}

	for retryCount := 10; retryCount > 0; retryCount-- {
		err = myq.Conn.Ping()
		if err == nil {
			// logger.Info("10", "Connect successfully to database.")
			break
		}
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		// err = logger.GetApplicationError(err).Init("12", "Connection failed to database.")
		return nil, err
	}

	h = myq

	return
}

func (h *SqlHandler) Exec(ctx context.Context, puery string, args ...interface{}) (res sql.Result, err error) {
	res, err = h.Conn.ExecContext(ctx, puery, args...)
	return
}

func (h *SqlHandler) Query(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	rows, err = h.Conn.QueryContext(ctx, query, args...)
	return
}

func (h *SqlHandler) QueryRow(ctx context.Context, query string, args ...interface{}) (row *sql.Row) {
	row = h.Conn.QueryRowContext(ctx, query, args...)
	return
}

func (h *SqlHandler) ExecWithTx(txFunc func(*sql.Tx) error) (err error) {
	tx, err := h.Conn.Begin()
	if err != nil {
		// err = logger.GetApplicationError(err).Init("xx", "Failed to start transaction.")
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// logger.Error(
			// 	_logger.NewApplicationError(p).
			// 		Init("xx", "An error has occured. Transaction is rolled back..."),
			// )
			err = tx.Rollback()
			panic(p)
		} else if err != nil {
			// logger.Error(
			// 	logger.GetApplicationError(err).
			// 		Init("xx", "An error has occured. Transaction is rolled back..."),
			// )
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = rollbackErr
			}
		} else {
			// logger.Debug("XX", "Begin transaction commit...")
			err = tx.Commit()
		}
	}()

	err = txFunc(tx)
	return
}
