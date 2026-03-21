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

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comments) (entity.Comments, error) {
	script := "INSERT INTO comments (email , comment) VALUES (?, ?)"
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
func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comments, error) {
	script := "SELECT id , email , comment FROM comments WHERE id = ?"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comments{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// tidak ada
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not found")
	}
}
func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comments, error) {
	script := "SELECT id , email , comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var comments []entity.Comments
	for rows.Next() {
		// ada
		comment := entity.Comments{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
