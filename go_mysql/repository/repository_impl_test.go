package repository

import (
	"context"
	"fmt"
	"go_mysql"
	"go_mysql/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(go_mysql.GetConnection())

	ctx := context.Background()
	newComment := entity.Comment{
		Email:   "repository@gmail.com",
		Comment: "Test Comment from Repository",
	}

	result, err := commentRepository.Insert(ctx, newComment)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result Insert Comment from Repository")
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(go_mysql.GetConnection())

	ctx := context.Background()
	comment, err := commentRepository.FindById(ctx, 72)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result Find By Id from Repository")
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(go_mysql.GetConnection())

	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result Find All from Repository")
	for _, comment := range comments {
		fmt.Println("Id: ", comment.Id, "| Email: ", comment.Email, "| Comment: ", comment.Comment)
	}
}

func TestUpdateById(t *testing.T) {
	commentRepository := NewCommentRepository(go_mysql.GetConnection())

	ctx := context.Background()
	updateComment := entity.Comment{
		Email:   "hhakimsetiawan@gmail.com",
		Comment: "Halo, apa kabar kawan?",
	}

	updatedComment, err := commentRepository.UpdateById(ctx, 2, updateComment)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result update comment by id")
	fmt.Println(updatedComment)
}
