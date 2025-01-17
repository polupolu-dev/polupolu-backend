package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/yourusername/poluback/internal/domain/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Get(ctx context.Context, id string) (*models.User, error) {
	query := `
		SELECT id, comment_ids, gender, age_group, occupation, political_view,
		       opinion_tone, speech_style, comment_length, background_knowledge,
		       emotion
		FROM users WHERE id = $1
	`

	var user models.User
	var commentIDsJSON []byte

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&commentIDsJSON,
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

	err = json.Unmarshal(commentIDsJSON, &user.CommentIDs)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	commentIDsJSON, err := json.Marshal(user.CommentIDs)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO users (id, comment_ids, gender, age_group, occupation,
		                  political_view, opinion_tone, speech_style,
		                  comment_length, background_knowledge, emotion)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err = r.db.ExecContext(ctx, query,
		user.ID,
		commentIDsJSON,
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

	return err
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	commentIDsJSON, err := json.Marshal(user.CommentIDs)
	if err != nil {
		return err
	}

	query := `
		UPDATE users 
		SET comment_ids = $2, gender = $3, age_group = $4,
		    occupation = $5, political_view = $6, opinion_tone = $7,
		    speech_style = $8, comment_length = $9,
		    background_knowledge = $10, emotion = $11
		WHERE id = $1
	`

	_, err = r.db.ExecContext(ctx, query,
		user.ID,
		commentIDsJSON,
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

	return err
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
