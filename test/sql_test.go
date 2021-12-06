package test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/WahidinAji/golang-database-mysql/database"
	"github.com/WahidinAji/golang-database-mysql/helpers"
	_ "github.com/go-sql-driver/mysql"
)

/** gunakan sql.Null`<tipedata>` untuk memberitahu bahwa column nya bisa null. contoh sql.NullString.
 *	httpsdan resultnya jadi struct, contoh data dri struct Customer
 *	{1 Wahidin {wahidin@gmail.com true} 200000 95 {1999-09-09 00:00:00 +0000 UTC true} true 2021-12-02 09:49:49 +0000 UTC}
 *	{2 Aji {aji@gmail.com true} 100000 85.5 {1998-08-08 00:00:00 +0000 UTC true} true 2021-12-02 09:49:49 +0000 UTC}
 *	{3 joko { false} 100000 87.5 {0001-01-01 00:00:00 +0000 UTC false} true 2021-12-02 11:04:47 +0000 UTC}
 */
type Customer struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Email sql.NullString `json:"email"`
	Balance int32 `json:"balance"`
	Rating float64 `json:"rating"`
	BirthDate sql.NullTime `json:"birth_date"`
	Married bool `json:"married"`
	CreatedAt time.Time `json:"created_at"`
}

/** test open connection*/
func TestOpenConnection(t *testing.T) {
	db := database.GetConnection()
	fmt.Println(db)
	db.Close()
}
/** test create data with ExecContext*/
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
/** test get data id and name only with QueryContext*/
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
/** test get data id, name, email, balance, rating, birth_date, married, created_at with QueryContext*/
func TestQuerySqlComplex(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customers"
	rows, err := db.QueryContext(ctx,script)
	helpers.PanicIfError(err)
	defer rows.Close()


	var result []Customer
	for rows.Next(){
		customers := Customer{}
		err := rows.Scan(&customers.Id, &customers.Name, &customers.Email, &customers.Balance, &customers.Rating, &customers.BirthDate, &customers.Married, &customers.CreatedAt)
		helpers.PanicIfError(err)
		// fmt.Println(customers)
		result = append(result, customers)
		/** valid buat ngecek, datanya null atau tidak. dan ini hanya isa digunakan ketika menggunakan sql.NullString atau sql.Null lainnya.
		  *	// fmt.Println("===========================")
		  *	// fmt.Println("id: ", customers.Id)
		  *	// fmt.Println("name: ", customers.Name)
		  *	// if customers.Email.Valid {
		  *	// 	fmt.Println("email: ", customers.Email)
		  *	// }
		  *	// fmt.Println("balance: ", customers.Balance)
		  *	// fmt.Println("rating: ", customers.Rating)
		  *	// if customers.BirthDate.Valid {
		  *	// 	fmt.Println("birth date: ", customers.BirthDate)
		  *	// }
		  *	// fmt.Println("married: ", customers.Married)
		  *	// fmt.Println("created at: ", customers.CreatedAt)
		*/
	}
	fmt.Println(result)
}

/** test delete all data with ExecContext*/
func TestQueryDelete(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()
	ctx := context.Background()

	querySql := "DELETE FROM customers"
	_, err := db.ExecContext(ctx, querySql)
	helpers.PanicIfError(err)
	fmt.Println("Success delete all customers")
} 

/** test get data id and name only with QueryContext if user login*/
func TestSqlinjection(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "salah"
	script := "SELECT username FROM users WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx,script,username,password)
	helpers.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		helpers.PanicIfError(err)
		fmt.Println("Login success ", username)
	} else {
		fmt.Println("Login failed")
	}
}

/** test get data id and name only with QueryContext if user login*/
func TestAutoIncrement(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	email := "ajisaka@gmail.com"
	comment := "Tes komen"	

	// script := "INSERT INTO comments(email, comment) VALUES(?, ?) ON DUPLICATE"
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	helpers.PanicIfError(err)
	insertid, err := result.LastInsertId()
	helpers.PanicIfError(err)
	fmt.Println("Success insert new comment with id: ", insertid)
}