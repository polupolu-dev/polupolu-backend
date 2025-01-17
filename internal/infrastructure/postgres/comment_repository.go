package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) GetByID(ctx context.Context, id string) ([]models.Comment, error) {
	query := `
		SELECT id, reply_to_id, user_id, content, created_at, 
		       feedback_scores, replied_ids
		FROM comments WHERE id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		var feedbackScoresJSON []byte
		var repliedIDsJSON []byte

		err := rows.Scan(
			&comment.ID,
			&comment.ReplyToID,
			&comment.UserID,
			&comment.Content,
			&comment.CreatedAt,
			&feedbackScoresJSON,
			&repliedIDsJSON,
		)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(feedbackScoresJSON, &comment.FeedbackScores)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(repliedIDsJSON, &comment.RepliedIDs)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (r *CommentRepository) GetByCommentID(ctx context.Context, commentID string) (*models.Comment, error) {
	query := `
		SELECT id, reply_to_id, user_id, content, created_at, 
		       feedback_scores, replied_ids
		FROM comments WHERE id = $1
	`

	var comment models.Comment
	var feedbackScoresJSON []byte
	var repliedIDsJSON []byte

	err := r.db.QueryRowContext(ctx, query, commentID).Scan(
		&comment.ID,
		&comment.ReplyToID,
		&comment.UserID,
		&comment.Content,
		&comment.CreatedAt,
		&feedbackScoresJSON,
		&repliedIDsJSON,
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(feedbackScoresJSON, &comment.FeedbackScores)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(repliedIDsJSON, &comment.RepliedIDs)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *CommentRepository) Create(ctx context.Context, comment *models.Comment) error {
	feedbackScoresJSON, err := json.Marshal(comment.FeedbackScores)
	if err != nil {
		return err
	}

	repliedIDsJSON, err := json.Marshal(comment.RepliedIDs)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO comments (id, reply_to_id, user_id, content, created_at, 
		                     feedback_scores, replied_ids)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err = r.db.ExecContext(ctx, query,
		comment.ID,
		comment.ReplyToID,
		comment.UserID,
		comment.Content,
		comment.CreatedAt,
		feedbackScoresJSON,
		repliedIDsJSON,
	)

	return err
}

func (r *CommentRepository) Update(ctx context.Context, comment *models.Comment) error {
	feedbackScoresJSON, err := json.Marshal(comment.FeedbackScores)
	if err != nil {
		return err
	}

	repliedIDsJSON, err := json.Marshal(comment.RepliedIDs)
	if err != nil {
		return err
	}

	query := `
		UPDATE comments 
		SET reply_to_id = $2, user_id = $3, content = $4, 
		    feedback_scores = $5, replied_ids = $6
		WHERE id = $1
	`

	_, err = r.db.ExecContext(ctx, query,
		comment.ID,
		comment.ReplyToID,
		comment.UserID,
		comment.Content,
		feedbackScoresJSON,
		repliedIDsJSON,
	)

	return err
}

func (r *CommentRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
