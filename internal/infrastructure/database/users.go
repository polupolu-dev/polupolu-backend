package database

import (
	"database/sql"
	"errors"

	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

// Create 新しいユーザーをデータベースに保存
func (r *UsersRepository) Create(user models.User) (*models.User, error) {
	query := `
        INSERT INTO users (id, comment_ids, gender, age_group, occupation, political_view, opinion_tone, speech_style, comment_length, background_knowledge, emotion)
		// VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := r.db.Exec(query,
		user.ID,
		user.CommentIDs,
		user.Gender,
		user.AgeGroup,
		user.Occupation,
		user.PoliticalView,
		user.OpinionTone,
		user.SpeechStyle,
		user.CommentLength,
		user.BackgroundKnowledge,
		user.Emotion,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Find ID に基づいて特定のユーザーを取得
func (r *UsersRepository) Find(id string) (*models.User, error) {
	query := `
        SELECT id, comment_ids, gender, age_group, occupation, political_view, opinion_tone, speech_style, comment_length, background_knowledge, emotion
		FROM users
		WHERE id = ?
	`
	row := r.db.QueryRow(query, id)

	var user models.User

	err := row.Scan(
		&user.ID,
		&user.CommentIDs,
		&user.Gender,
		&user.AgeGroup,
		&user.Occupation,
		&user.PoliticalView,
		&user.OpinionTone,
		&user.SpeechStyle,
		&user.CommentLength,
		&user.BackgroundKnowledge,
		&user.Emotion,
		&user.ID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// FindList 指定された ID のユーザーを取得
func (r *UsersRepository) FindList(id string) ([]models.User, error) {
	query := `
        SELECT id, comment_ids, gender, age_group, occupation, political_view, opinion_tone, speech_style, comment_length, background_knowledge, emotion
		FROM users
		WHERE id = ?
	`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.CommentIDs,
			&user.Gender,
			&user.AgeGroup,
			&user.Occupation,
			&user.PoliticalView,
			&user.OpinionTone,
			&user.SpeechStyle,
			&user.CommentLength,
			&user.BackgroundKnowledge,
			&user.Emotion,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Update 特定のユーザーを更新
func (r *UsersRepository) Update(user models.User) (*models.User, error) {
	query := `
		UPDATE users
		SET comment_ids = ?, gender = ?, age_group = ?, occupation = ?, political_view = ?, opinion_tone = ?, speech_style = ?, comment_length = ?, background_knowledge = ?, emotion = ?
		WHERE id = ?
	`
	_, err := r.db.Exec(query,
		user.CommentIDs,
		user.Gender,
		user.AgeGroup,
		user.Occupation,
		user.PoliticalView,
		user.OpinionTone,
		user.SpeechStyle,
		user.CommentLength,
		user.BackgroundKnowledge,
		user.Emotion,
		user.ID,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Delete 指定された ID のユーザーを削除
func (r *UsersRepository) Delete(id string) error {
	query := `DELETE FROM users WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
