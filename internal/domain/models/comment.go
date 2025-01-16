package models

import (
	"errors"
	"time"
)

type Comment struct {
	ID             string         `json:"id"`
	ReplyToID      string         `json:"reply_to_id"`
	RepliedIDs     []string       `json:"replied_ids"`
	UserID         string         `json:"user_id"`
	Content        string         `json:"content"`
	CreatedAt      time.Time      `json:"created_at"`
	FeedbackScores FeedbackScores `json:"feedback_scores"`
}

func (n *News) CommentValidate() error {
	if n.ID == "" {
		return errors.New("ID is required")
	}

	return nil
}
