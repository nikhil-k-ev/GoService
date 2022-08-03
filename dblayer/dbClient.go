package dblayer

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBRepo struct {
	SqlDB *sql.DB
}

func (dc *DBRepo) DBConnect() error {
	cnnStr := "Write Connection String here"

	fmt.Print("\n\n\n", cnnStr, "\n\n\n")

	var err error

	dc.SqlDB, err = sql.Open("postgres", cnnStr)
	if err != nil {
		fmt.Print("\n\n\n", err, "\n\n\n")
		return err
	}

	err = dc.SqlDB.Ping()
	if err != nil {
		fmt.Print("\n\n\n", err, "\n\n\n")
		return err
	}

	return nil

}
