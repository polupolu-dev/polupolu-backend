package postgresql

import (
	"database/sql"
	"errors"

	"github.com/go-playground/validator"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type commentsRepositoryImpl struct {
	*PostgresqlDependency
}

func NewCommentsRepository(deps *PostgresqlDependency) (*commentsRepositoryImpl, error) {
	err := validator.New().Struct(deps)
	if err != nil {
		return &commentsRepositoryImpl{}, err
	}

	impl := commentsRepositoryImpl{
		deps,
	}

	return &impl, nil
}

// Create 新しいコメントをデータベースに保存
func (r *commentsRepositoryImpl) Create(comment models.Comment) (*models.Comment, error) {
	query := `
    	INSERT INTO comments (id, reply_to_id, replied_ids, user_id, content, created_at, empathy, insight, mediocre)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(query,
		comment.ID,
		comment.ReplyToID,
		comment.RepliedIDs,
		comment.UserID,
		comment.Content,
		comment.CreatedAt,
		comment.FeedbackScores.Empathy,
		comment.FeedbackScores.Insight,
		comment.FeedbackScores.Mediocre,
	)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// Find ID に基づいて特定のコメントを取得
func (r *commentsRepositoryImpl) Find(id string) (*models.Comment, error) {
	query := `
		SELECT id, reply_to_id, replied_ids, user_id, content, created_at, empathy, insight, mediocre
		FROM comments
		WHERE id = ?
	`
	row := r.db.QueryRow(query, id)

	var comment models.Comment

	err := row.Scan(
		&comment.ID,
		&comment.ReplyToID,
		&comment.RepliedIDs,
		&comment.UserID,
		&comment.Content,
		&comment.CreatedAt,
		&comment.FeedbackScores.Empathy,
		&comment.FeedbackScores.Insight,
		&comment.FeedbackScores.Mediocre,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("comment not found")
		}
		return nil, err
	}

	return &comment, nil
}

// FindList 指定された ID（例えば news_id または user_id）に関連するコメント一覧を取得
func (r *commentsRepositoryImpl) FindList(id string) ([]models.Comment, error) {
	query := `
		SELECT id, reply_to_id, replied_ids, user_id, content, created_at, empathy, insight, mediocre
		FROM comments
		WHERE reply_to_id = ? OR user_id = ?
	`
	rows, err := r.db.Query(query, id, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment

		err := rows.Scan(
			&comment.ID,
			&comment.ReplyToID,
			&comment.RepliedIDs,
			&comment.UserID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.FeedbackScores.Empathy,
			&comment.FeedbackScores.Insight,
			&comment.FeedbackScores.Mediocre,
		)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

// Update 特定のコメントを更新
func (r *commentsRepositoryImpl) Update(comment models.Comment) (*models.Comment, error) {
	query := `
		UPDATE comments
		SET replied_ids = ?, content = ?, empathy = ?, insight = ?, mediocre = ?
		WHERE id = ?
	`
	_, err := r.db.Exec(query,
		comment.RepliedIDs,
		comment.Content,
		comment.FeedbackScores.Empathy,
		comment.FeedbackScores.Insight,
		comment.FeedbackScores.Mediocre,
		comment.ID,
	)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// Delete 指定された ID のコメントを削除
func (r *commentsRepositoryImpl) Delete(id string) error {
	query := `DELETE FROM comments WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("comment not found")
	}
	return nil
}
