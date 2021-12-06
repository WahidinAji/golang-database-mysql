package test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/WahidinAji/golang-database-mysql/database"
	"github.com/WahidinAji/golang-database-mysql/helpers"
)

func TestTransaction(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	helpers.PanicIfError(err)

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	// do transactions
	for i := 0; i < 10; i++ {
		email := "aji" + strconv.Itoa(i) + "@gmail.com";
		comment := "Comment to " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		helpers.PanicIfError(err)
		id, err :=result.LastInsertId()
		helpers.PanicIfError(err)
		fmt.Println("Comment Id", id)
	}
	err = tx.Commit()
	// err = tx.Rollback()
	helpers.PanicIfError(err) //if transaction fail, return error
}