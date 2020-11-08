package infrastructure

import (
	"database/sql"
	"fmt"

	interfaces "github.com/Maser-DC/Interfaces"
	_ "github.com/lib/pq"
)

type PostgresHandler struct {
	Conn *sql.DB
}

func (handler *PostgresHandler) Close() error {
	err := handler.Conn.Close()
	return err
}
func (handler *PostgresHandler) Execute(statement string) error {
	_, err := handler.Conn.Exec(statement)

	return err
}
func (handler *PostgresHandler) Query(statement string) (interfaces.Row, error) {
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(PostgresRow), err
	}
	row := new(PostgresRow)
	row.Rows = rows
	return row, err
}

type PostgresRow struct {
	Rows *sql.Rows
}

func (r PostgresRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	return err
}
func (r PostgresRow) Next() bool {
	return r.Rows.Next()
}
func NewPostgresHandler(host string, port int, user string, password string, dbname string) *PostgresHandler {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = conn.Ping()
	if err != nil {
		panic(err)
	}
	postgresHandler := new(PostgresHandler)
	postgresHandler.Conn = conn

	fmt.Println("Successfully connected!")

	return postgresHandler
}
