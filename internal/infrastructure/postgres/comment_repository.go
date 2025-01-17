package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) GetByID(ctx context.Context, id uuid.UUID) ([]models.Comment, error) {
	query := `
		SELECT id, reply_to_id, user_id, content, created_at, 
	           empathy, insight, mediocre, replied_ids
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

		err := rows.Scan(
			&comment.ID,
			&comment.ReplyToID,
			&comment.UserID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.FeedbackScores.Empathy,
			&comment.FeedbackScores.Insight,
			&comment.FeedbackScores.Mediocre,
			&comment.RepliedIDs,
		)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (r *CommentRepository) GetByCommentID(ctx context.Context, commentID uuid.UUID) (*models.Comment, error) {
	query := `
		SELECT id, reply_to_id, user_id, content, created_at, 
		       empathy, insight, mediocre, replied_ids
		FROM comments WHERE id = $1
	`

	var comment models.Comment

	err := r.db.QueryRowContext(ctx, query, commentID).Scan(
		&comment.ID,
		&comment.ReplyToID,
		&comment.UserID,
		&comment.Content,
		&comment.CreatedAt,
		&comment.FeedbackScores.Empathy,
		&comment.FeedbackScores.Insight,
		&comment.FeedbackScores.Mediocre,
		&comment.RepliedIDs,
	)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *CommentRepository) Create(ctx context.Context, comment *models.Comment) error {
	query := `
		INSERT INTO comments (id, reply_to_id, user_id, content, created_at, 
		                     empathy, insight, mediocre, replied_ids)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.ExecContext(ctx, query,
		comment.ID,
		comment.ReplyToID,
		comment.UserID,
		comment.Content,
		comment.CreatedAt,
		comment.FeedbackScores.Empathy,
		comment.FeedbackScores.Insight,
		comment.FeedbackScores.Mediocre,
		comment.RepliedIDs,
	)

	return err
}

func (r *CommentRepository) Update(ctx context.Context, comment *models.Comment) error {
	query := `
		UPDATE comments 
		SET reply_to_id = $2, user_id = $3, content = $4, 
            empathy = $5, insight = $6, mediocre = $7, 
		    feedback_scores = $8, replied_ids = $9
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query,
		comment.ID,
		comment.ReplyToID,
		comment.UserID,
		comment.Content,
		comment.FeedbackScores.Empathy,
		comment.FeedbackScores.Insight,
		comment.FeedbackScores.Mediocre,
		comment.RepliedIDs,
	)

	return err
}

func (r *CommentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
