package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	fmt.Println(IsDBUp("OTA-DB", "mysql", "hotelurb_dba:QHXveCUdpcppy-2j@tcp(database_ota.hud:3306)/ota"))
	fmt.Println(IsDBUp("NHU-DB", "mysql", "hotelurb_dba:QHXveCUdpcppy-2j@tcp(database_nhu.hud:3306)/novo_hu"))
}

// IsDBUp ping dependency.
func IsDBUp(DBName, DBDriver, DBDSN string) bool {
	fmt.Println(DBName, DBDriver, DBDSN)
	db, err := sqlx.Connect(DBDriver, DBDSN)
	if err != nil {
		fmt.Println(
			"Ping to %s fail.",
			DBName,
		)
		return false
	}
	db.Close()
	return true
}
