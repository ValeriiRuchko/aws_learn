package db_ops

import (
	"database/sql"
	"fmt"
	"os"
)

func Initialize_Connection() *sql.DB {
	connectionUrl := os.Getenv("TursoDatabaseURL") + "?authToken=" + os.Getenv("TursoAuthToken")

	db, err := sql.Open("libsql", connectionUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", connectionUrl, err)
		os.Exit(1)
	}
	return db
}
