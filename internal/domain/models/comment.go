package models

import (
	"time"
)

type Comment struct {
	ID             string    `json:"id"`
	ReplyToID      string    `json:"reply_to_id"`
	RepliedIDs     []string  `json:"replied_ids"`
	UserID         string    `json:"user_id"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
	FeedbackScores struct {
		Empathy  uint32 `json:"empathy"`
		Insight  uint32 `json:"insight"`
		Mediocre uint32 `json:"mediocre"`
	} `json:"feedback_scores"`
}
