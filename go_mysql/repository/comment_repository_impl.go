package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go_mysql/entity"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	var resultComment entity.Comment
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"

	rows, err := repository.DB.QueryContext(ctx, script, id)
	if err != nil {
		return resultComment, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&resultComment.Id, &resultComment.Email, &resultComment.Comment)
		if err != nil {
			return resultComment, err
		}

		return resultComment, nil
	} else {
		return resultComment, errors.New("Comment dengan Id " + strconv.Itoa(int(id)) + " tidak ada")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"

	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		var comment entity.Comment

		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (repository *commentRepositoryImpl) UpdateById(ctx context.Context, id int32, comment entity.Comment) (entity.Comment, error) {
	updatedComment := entity.Comment{}
	script := "UPDATE comments SET email = ?, comment = ? WHERE id = ?"

	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment, id)
	if err != nil {
		return updatedComment, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return updatedComment, err
	}

	updatedComment.Id = id
	updatedComment.Email = comment.Email
	updatedComment.Comment = comment.Comment

	fmt.Println("Result", result)
	fmt.Println("Result rows affected", strconv.Itoa(int(rowsAffected)))

	return updatedComment, nil
}

func (repository *commentRepositoryImpl) DeleteById(ctx context.Context, id int32) error {
	return nil
}
