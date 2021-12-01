package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/WahidinAji/golang-database-mysql/database"
	"github.com/WahidinAji/golang-database-mysql/helpers"
	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db := database.GetConnection()
	fmt.Println(db)
	db.Close()
}

func TestExecSql(t *testing.T) {
	db := database.GetConnection()
	fmt.Println(db)
	defer db.Close()

	ctx := context.Background()

	querySql := "INSERT INTO customers(name) VALUES(?)"
	_, err := db.ExecContext(ctx, querySql,"Aji Saka ")
	helpers.PanicIfError(err)
	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customers"
	rows, err := db.QueryContext(ctx,script)
	helpers.PanicIfError(err)
	defer rows.Close()

	for rows.Next(){
		var id int
		var name string
		err := rows.Scan(&id, &name)
		helpers.PanicIfError(err)
		fmt.Println("id: ", id)
		fmt.Println("name: ", name)
	}
}

func TestQueryDelete(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()
	ctx := context.Background()

	querySql := "DELETE FROM customers"
	_, err := db.ExecContext(ctx, querySql)
	helpers.PanicIfError(err)
	fmt.Println("Success delete all customers")
} 
