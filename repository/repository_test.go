package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/WahidinAji/golang-database-mysql/database"
	"github.com/WahidinAji/golang-database-mysql/entity"
	"github.com/WahidinAji/golang-database-mysql/helpers"
)

func TestCommentInserter(t *testing.T){
	commentRepository := NewCommentRepository(database.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email: "repository@test.com",
		Comment: "Test Repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	helpers.PanicIfError(err) 
	fmt.Println(result)
}

func TestFindByID(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(),21)
	helpers.PanicIfError(err)
	fmt.Println(comment)
}
func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())

	comments, err := commentRepository.FindByAll(context.Background())
	helpers.PanicIfError(err)
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
