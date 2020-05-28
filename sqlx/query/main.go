package main

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := buildcfg()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	otaDB, err := newDatabase(ctx, "mysql", cfg.dsn)
	if err != nil {
		logrus.WithError(err).Fatal("Error in load database instance.")
	}
	defer otaDB.Connection.Close()

	for _, t := range []string{
		fmt.Sprintf("%s.%s", cfg.dbName, cfg.table),
		cfg.table,
	} {
		query := fmt.Sprintf(cfg.query, t, cfg.id)
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

	return &database{Connection: conn, Name: ""}, nil
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
