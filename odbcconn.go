package main

import (
	"github.com/alexbrainman/odbc"
)

func main() {
	conn, _ := odbc.Connect("ODBC","DSN=impala")
	//conn, _ := sql.Open("ODBC","DSN=impala")
	}
	fmt.Println("connected")
	conn.Close()
}
