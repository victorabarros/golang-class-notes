package main

import (
    _ "github.com/go-sql-driver/mysql"
    "context"
    "github.com/jmoiron/sqlx"
    "log"
)

const (
    databaseDriver = "mysql"
    databaseDSN    = "user:password@protocol(ip:port)/schema"
)

type Sku struct {
    SkuType string `db:"sku_type"`
    SkuCode  string `db:"sku_code"`
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    db, err := sqlx.ConnectContext(ctx, databaseDriver, databaseDSN)
    if err != nil {
        log.Println(err)
        return
    }
    skus := []Sku{}
    db.Select(&skus, "SELECT sku_type, sku_code FROM hotel")
    log.Println("Skus")
    for _, sku := range skus{
        log.Println(sku.SkuType, sku.SkuCode)
    }

    // rows, err := db.Query(`SELECT sku_type FROM hotel`)
    // if err != nil {
    //     log.Println(err)
    //     return
    // }

    // defer rows.Close()
    // if !rows.Next() {
    //     log.Println(err)
    //     return
    // }
    // log.Println(rows)
}