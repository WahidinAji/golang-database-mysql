package test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/WahidinAji/golang-database-mysql/database"
	"github.com/WahidinAji/golang-database-mysql/helpers"
)

func TestPrepareStatement(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	stmt, err := db.PrepareContext(ctx,script)
	helpers.PanicIfError(err)
	defer stmt.Close() 

	for i := 0; i < 10; i++ {
		email := "aji" + strconv.Itoa(i) + "@gmail.com";
		comment := "Comment to " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx,email, comment)
		helpers.PanicIfError(err)
		id, err :=result.LastInsertId()
		helpers.PanicIfError(err)
		fmt.Println("Comment Id", id)
	}
}