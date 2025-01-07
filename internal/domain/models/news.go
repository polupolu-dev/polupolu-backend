package models

import (
	"time"
)

type News struct {
	ID             string    `json:"id"`
	Category       string    `json:"category"`
	Title          string    `json:"title"`
	Source         string    `json:"source"`
	URL            string    `json:"url"`
	Summary        string    `json:"summary"`
	PublishedAt    time.Time `json:"publishedAt"`
	FeedbackScores struct {
		Empathy  uint32 `json:"empathy"`
		Insight  uint32 `json:"insight"`
		Mediocre uint32 `json:"mediocre"`
	}
	CommentIDs []string `json:"commentIDs"`
}
