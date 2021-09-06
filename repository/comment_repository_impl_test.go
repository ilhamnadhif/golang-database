package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "Muhammad@tet.com",
		Comment: "Test Repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	comment, err := commentRepository.FindById(ctx, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestCommendFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println("=======================================================")
		fmt.Println(comment)
	}
}

func TestCommentUpdate(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "Halooooo@tet.com",
		Comment: "Test Repository",
	}
	result, err := commentRepository.UpdateById(ctx, 36, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentDelete(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	result, err := commentRepository.DeleteById(ctx, 35)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
