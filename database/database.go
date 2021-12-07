package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/WahidinAji/golang-database-mysql/config"
	_ "github.com/go-sql-driver/mysql"
)

/** database */

func GetConnection() *sql.DB {
	get := config.GetEnvWithKey
	user := get("DB_USER")
	pass := get("DB_PASS")
	host := get("DB_HOST")
	port := get("DB_PORT")
	dbname := get("DB_NAME")
	//parseTime=true digunakan untuk mengkonversi []uint dari db ke dateTime string. jadi si golang tidak perlu lagi convert manual.
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",user,pass,host,port,dbname)
	db, err := sql.Open("mysql",dsn)
	if err != nil {
		fmt.Println("Connection Failed!!!")
		panic(err)
	}
	db.SetMaxIdleConns(10) //min 10 connection
	db.SetMaxOpenConns(100) //max 100 connection
	db.SetConnMaxIdleTime(5 * time.Minute) //if iin 5 minutes nothing happen? db will close the connection
	db.SetConnMaxLifetime(60 * time.Minute) //after 60 minutes, the connection will create new connection
	return db
}