package main

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	logLevel   = "INFO"
	dbUser     = "dbUser"
	dbPassword = "dbPassword"
	dbHost     = "dbHost"
	dbName     = "dbName"
)

var (
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s",
		dbUser,
		dbPassword,
		dbHost,
		dbName)
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	otaDB, err := newDatabase(ctx, "mysql", dsn)
	if err != nil {
		logrus.WithError(err).Fatal("Error in load database instance.")
	}
	defer otaDB.Connection.Close()

	table := "table"
	query := "SELECT DISTINCT id FROM %s WHERE id IN (%s);"

	for _, t := range []string{
		fmt.Sprintf("%s.%s", dbName, table),
		table,
	} {
		query := fmt.Sprintf(query, t, "id")
		fmt.Println(otaDB.selectHotelIDs(query))
	}
}

type database struct {
	Connection *sqlx.DB
	Name       string
}

func newDatabase(ctx context.Context, driver string, dsn string) (*database, error) {
	conn, err := sqlx.ConnectContext(ctx, driver, dsn)
	if err != nil {
		return nil, err
	}

	return &database{Connection: conn, Name: dbName}, nil
}

func (db *database) selectHotelIDs(query string) ([]int, error) {
	var ids []int
	logrus.Infof("Querying on %s-DB: %s", db.Name, query)
	err := db.Connection.Select(&ids, query)
	if err != nil {
		return nil, err
	}
	return ids, err
}
