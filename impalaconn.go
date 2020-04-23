package main

import (
	"fmt"
	_ "github.com/alexbrainman/odbc"
	"github.com/jmoiron/sqlx"
)

type Res struct {
	Nam  string `db:"name"`
	Comm string `db:"comment"`
}

func main() {
	conn, err := sqlx.Connect("odbc", "DSN=IMPALA")
	if err != nil {
		fmt.Println("ERROR", err)
		panic("could not connect")
	}

	var res []Res
	query := "SELECT name, comment FROM testdb.t1"
	if err := conn.Select(&res, query); err != nil {
		fmt.Println("ERROR", err)
		panic("error running query")
	}

	fmt.Println(res)
}
