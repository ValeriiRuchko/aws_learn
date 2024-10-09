package db_ops

import (
	"database/sql"
	"fmt"
	"os"
)

func Initialize_Connection() *sql.DB {
	url := os.Getenv("TursoDatabaseURL") + "?authToken=" + os.Getenv("TursoAuthToken")

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	return db
}
