package golang_database_mysql

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/WahidinAji/golang-database-mysql/utils"
	_ "github.com/go-sql-driver/mysql"
)


func TestOpenConnection(t *testing.T) {
	get := utils.GetEnvWithKey
	user := get("DB_USER")
	pass := get("DB_PASS")
	host := get("DB_HOST")
	port := get("DB_PORT")
	dbname := get("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",user,pass,host,port,dbname)
	fmt.Println(dsn)
	// dsn := "root:@tcp(localhost:3306)/go_database"
	db, err := sql.Open("mysql",dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}