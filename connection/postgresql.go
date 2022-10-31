package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func DatabaseConnect() {

	// TEMPLATE "postgres://username:password@localhost:5432/databasename"
	databaseUrl := "postgres://postgres:766880@localhost:5432/db_project"
	
	var err error
	Conn, err = pgx.Connect(context.Background(), databaseUrl)
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Success connect to database")
	}
}
