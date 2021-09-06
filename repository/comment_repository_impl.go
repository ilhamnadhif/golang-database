package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"
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
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	var comment entity.Comment
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// nggak ada
		return comment, errors.New(("Id " + strconv.Itoa(int(id)) + " Not Found"))
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
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
func (repository *commentRepositoryImpl) UpdateById(ctx context.Context, id int32, comment entity.Comment) (entity.Comment, error) {
	//tx, err := repository.DB.Begin()
	//if err != nil {
	//	return comment, err
	//}
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return comment, err
	}
	if rows.Next() {
		// ada
		script := "UPDATE comments SET email = ?, comment = ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment, id)
		if err != nil {
			return comment, err
		}
		comment.Id = id
		return comment, nil
	} else {
		// nggak ada
		return comment, errors.New(("Id " + strconv.Itoa(int(id)) + " Not Found"))
	}
}

func (repository *commentRepositoryImpl) DeleteById(ctx context.Context, id int32) (string, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return "gagal dihapus", err
	}
	if rows.Next() {
		// ada
		script := "DELETE FROM comments WHERE id = ?"
		_, erro := repository.DB.ExecContext(ctx, script, id)
		if erro != nil {
			return "Id " + strconv.Itoa(int(id)) + " Gagal Dihapus", erro
		}
		return "Id " + strconv.Itoa(int(id)) + " Berhasil Dihapus", nil
	} else {
		// nggak ada
		return "gagal dihapus", errors.New(("Id " + strconv.Itoa(int(id)) + " Not Found"))
	}
}
