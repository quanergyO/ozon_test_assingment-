package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/quanergyo/ozon-test-assingment/graph/model"
	"strconv"
)

type Comment struct {
	db *sqlx.DB
}

func NewComment(db *sqlx.DB) *Comment {
	return &Comment{
		db: db,
	}
}

func (r *Comment) CreateComment(userId int, postID int, parentID *string, content string) (*model.Comment, error) {
	// TODO check is open comments and is postID exists
	var commentId int
	var rows string
	var insertPlacement string
	args := make([]interface{}, 0, 4)
	args = append(args, userId, postID)
	if parentID != nil {
		rows = "userId, postId, parentId, content"
		insertPlacement = "$1, $2, $3, $4"
		args = append(args, *parentID)
	} else {
		rows = "userId, postId, content"
		insertPlacement = "$1, $2, $3"
	}
	args = append(args, content)

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s) RETURNING ID", commentsTable, rows, insertPlacement)
	row := r.db.QueryRow(query, args...)
	if err := row.Scan(&commentId); err != nil {
		return nil, err
	}

	return &model.Comment{
		ID:       strconv.Itoa(commentId),
		UserID:   strconv.Itoa(userId),
		PostID:   strconv.Itoa(postID),
		ParentID: parentID,
		Content:  content,
	}, nil
}

func (r *Comment) GetAllComments(postId int) ([]*model.Comment, error) {
	comments := make([]*model.Comment, 0)
	rows := "id, userId, postId, parentId, content"
	query := fmt.Sprintf("SELECT %s FROM %s WHERE postId=%d", rows, commentsTable, postId)
	if err := r.db.Select(&comments, query); err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *Comment) GetCommentsByPage(postId int, page int) ([]*model.Comment, error) {
	comments := make([]*model.Comment, 0)
	offset := calculateOffset(page)
	rows := "id, userId, postId, parentId, content"
	query := fmt.Sprintf("SELECT %s FROM %s WHERE postId=%d LIMIT %d OFFSET %d", rows, commentsTable, postId, commentsPerPage, offset)
	if err := r.db.Select(&comments, query); err != nil {
		return nil, err
	}

	return comments, nil
}

func calculateOffset(page int) int {
	return page * commentsPerPage
}
