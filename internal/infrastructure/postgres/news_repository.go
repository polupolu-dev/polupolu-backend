package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type NewsRepository struct {
	db *sql.DB
}

func NewNewsRepository(db *sql.DB) *NewsRepository {
	return &NewsRepository{db: db}
}

func (r *NewsRepository) GetAll(ctx context.Context) ([]models.News, error) {
	query := `
		SELECT id, category, title, source, url, summary, published_at,
		       feedback_scores, comment_ids
		FROM news
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var newsList []models.News
	for rows.Next() {
		var news models.News
		var feedbackScoresJSON []byte
		var commentIDsJSON []byte

		err := rows.Scan(
			&news.ID,
			&news.Category,
			&news.Title,
			&news.Source,
			&news.URL,
			&news.Summary,
			&news.PublishedAt,
			&feedbackScoresJSON,
			&commentIDsJSON,
		)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(feedbackScoresJSON, &news.FeedbackScores)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(commentIDsJSON, &news.CommentIDs)
		if err != nil {
			return nil, err
		}

		newsList = append(newsList, news)
	}

	return newsList, nil
}

func (r *NewsRepository) GetByID(ctx context.Context, id string) (*models.News, error) {
	query := `
		SELECT id, category, title, source, url, summary, published_at,
		       feedback_scores, comment_ids
		FROM news WHERE id = $1
	`

	var news models.News
	var feedbackScoresJSON []byte
	var commentIDsJSON []byte

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&news.ID,
		&news.Category,
		&news.Title,
		&news.Source,
		&news.URL,
		&news.Summary,
		&news.PublishedAt,
		&feedbackScoresJSON,
		&commentIDsJSON,
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(feedbackScoresJSON, &news.FeedbackScores)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(commentIDsJSON, &news.CommentIDs)
	if err != nil {
		return nil, err
	}

	return &news, nil
}

func (r *NewsRepository) GetByCategory(ctx context.Context, category string) ([]models.News, error) {
	query := `
		SELECT id, category, title, source, url, summary, published_at,
		       feedback_scores, comment_ids
		FROM news WHERE category = $1
	`
	rows, err := r.db.QueryContext(ctx, query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var newsList []models.News
	for rows.Next() {
		var news models.News
		var feedbackScoresJSON []byte
		var commentIDsJSON []byte

		err := rows.Scan(
			&news.ID,
			&news.Category,
			&news.Title,
			&news.Source,
			&news.URL,
			&news.Summary,
			&news.PublishedAt,
			&feedbackScoresJSON,
			&commentIDsJSON,
		)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(feedbackScoresJSON, &news.FeedbackScores)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(commentIDsJSON, &news.CommentIDs)
		if err != nil {
			return nil, err
		}

		newsList = append(newsList, news)
	}

	return newsList, nil
}

func (r *NewsRepository) Create(ctx context.Context, news *models.News) error {
	feedbackScoresJSON, err := json.Marshal(news.FeedbackScores)
	if err != nil {
		return err
	}

	commentIDsJSON, err := json.Marshal(news.CommentIDs)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO news (id, category, title, source, url, summary, published_at,
		                 feedback_scores, comment_ids)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err = r.db.ExecContext(ctx, query,
		news.ID,
		news.Category,
		news.Title,
		news.Source,
		news.URL,
		news.Summary,
		news.PublishedAt,
		feedbackScoresJSON,
		commentIDsJSON,
	)

	return err
}

func (r *NewsRepository) Update(ctx context.Context, news *models.News) error {
	feedbackScoresJSON, err := json.Marshal(news.FeedbackScores)
	if err != nil {
		return err
	}

	commentIDsJSON, err := json.Marshal(news.CommentIDs)
	if err != nil {
		return err
	}

	query := `
		UPDATE news 
		SET category = $2, title = $3, source = $4, url = $5,
		    summary = $6, published_at = $7, feedback_scores = $8,
		    comment_ids = $9
		WHERE id = $1
	`

	_, err = r.db.ExecContext(ctx, query,
		news.ID,
		news.Category,
		news.Title,
		news.Source,
		news.URL,
		news.Summary,
		news.PublishedAt,
		feedbackScoresJSON,
		commentIDsJSON,
	)

	return err
}

func (r *NewsRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM news WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
