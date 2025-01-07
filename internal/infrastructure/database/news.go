package database

import (
	"database/sql"
	"errors"

	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type NewsRepository struct {
	db *sql.DB
}

func NewNewsRepository(db *sql.DB) *NewsRepository {
	return &NewsRepository{db: db}
}

// Create 新しいニュースをデータベースに保存
func (r *NewsRepository) Create(news models.News) (*models.News, error) {
	query := `
    	INSERT INTO news (id, category, title, source, url, summary, empathy, publishedAt, empathy, insight, mediocre)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(query,
		news.ID,
		news.Category,
		news.Title,
		news.Source,
		news.URL,
		news.Summary,
		news.PublishedAt,
		news.FeedbackScores.Empathy,
		news.FeedbackScores.Insight,
		news.FeedbackScores.Mediocre,
	)
	if err != nil {
		return nil, err
	}
	return &news, nil
}

// Find ID に基づいて特定のニュースを取得
func (r *NewsRepository) Find(id string) (*models.News, error) {
	query := `
		SELECT id, category, title, source, url, summary, empathy, publishedAt, empathy, insight, mediocre
		FROM news
		WHERE id = ?
	`
	row := r.db.QueryRow(query, id)

	var news models.News

	err := row.Scan(
		&news.ID,
		&news.Category,
		&news.Title,
		&news.Source,
		&news.URL,
		&news.Summary,
		&news.PublishedAt,
		&news.FeedbackScores.Empathy,
		&news.FeedbackScores.Insight,
		&news.FeedbackScores.Mediocre,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("news not found")
		}
		return nil, err
	}

	return &news, nil
}

// FindAll すべてのニュース取得
func (r *NewsRepository) FindAll(category string) ([]models.News, error) {
	query := `
        SELECT id, category, title, source, url, summary, empathy, publishedAt, empathy, insight, mediocre
        FROM news
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var news []models.News
	for rows.Next() {
		var aNews models.News

		err := rows.Scan(
			&aNews.ID,
			&aNews.Category,
			&aNews.Title,
			&aNews.Source,
			&aNews.URL,
			&aNews.Summary,
			&aNews.PublishedAt,
			&aNews.FeedbackScores.Empathy,
			&aNews.FeedbackScores.Insight,
			&aNews.FeedbackScores.Mediocre,
		)
		if err != nil {
			return nil, err
		}

		news = append(news, aNews)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return news, nil
}

// FindList 指定されたカテゴリのニュース一覧を取得
func (r *NewsRepository) FindList(category string) ([]models.News, error) {
	query := `
		SELECT id, category, title, source, url, summary, empathy, publishedAt, empathy, insight, mediocre
		FROM news
		WHERE category = ?
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var news []models.News
	for rows.Next() {
		var aNews models.News

		err := rows.Scan(
			&aNews.ID,
			&aNews.Category,
			&aNews.Title,
			&aNews.Source,
			&aNews.URL,
			&aNews.Summary,
			&aNews.PublishedAt,
			&aNews.FeedbackScores.Empathy,
			&aNews.FeedbackScores.Insight,
			&aNews.FeedbackScores.Mediocre,
			&aNews.Category,
		)
		if err != nil {
			return nil, err
		}

		news = append(news, aNews)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return news, nil
}

// Update 特定のニュースを更新
func (r *NewsRepository) Update(news models.News) (*models.News, error) {
	query := `
		UPDATE news
		SET category = ?, title = ?, source = ?, url = ?, summary = ?, empathy = ?, insight = ?, mediocre = ?
		WHERE id = ?
	`
	_, err := r.db.Exec(query,
		news.Category,
		news.Title,
		news.Source,
		news.URL,
		news.Summary,
		news.FeedbackScores.Empathy,
		news.FeedbackScores.Insight,
		news.FeedbackScores.Mediocre,
		news.ID,
	)
	if err != nil {
		return nil, err
	}
	return &news, nil
}

// Delete 指定された ID のニュースを削除
func (r *NewsRepository) Delete(id string) error {
	query := `DELETE FROM news WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("news not found")
	}
	return nil
}
