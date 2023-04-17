package postgresSQL

import (
	"database/sql"
	"github.com/jdcd/account_balance/pkg"
	"time"

	_ "github.com/lib/pq"
)

func GetConnection(connectionData string) *sql.DB {
	db, err := sql.Open("postgres", connectionData)
	if err != nil {
		pkg.ErrorLogger().Printf("error connecting to database: %s", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		pkg.ErrorLogger().Printf("error making ping to database: %s", err)
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	return db
}
