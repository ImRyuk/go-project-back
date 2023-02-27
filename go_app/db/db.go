package db

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

func DbConnection() (*sql.DB, error) {
	db_conf := os.Getenv("MARIADB_USER")+ ":" + os.Getenv("MARIADB_ROOT_PASSWORD") + "@tcp(" + os.Getenv("MARIADB_HOST") + ":3306)/" + os.Getenv("MARIADB_DATABASE")
	fmt.Println(db_conf)
	db, err := sql.Open("mysql", db_conf)

	if err != nil {
		return nil, err
	}
	return db, nil
}
