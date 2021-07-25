package dao

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

var Db *sql.DB

type Customer struct {
	CustomerId string
	Name       string
}

type code interface {
	NotFound() error
}

func GetResult() (Customer, error) {
	var customer Customer

	Db, err := sql.Open("mysql", "root:ttalbe@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer Db.Close()
	sql := "select * from test"
	row := Db.QueryRow(sql)
	err = row.Scan(&customer.CustomerId, &customer.Name)
	return customer, errors.Wrapf(err, fmt.Sprintf("sql: %s error: %v", sql, err))

}
