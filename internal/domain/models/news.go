package models

import (
	"time"
)

type FeedbackScores struct {
	Empathy  uint32 `json:"empathy"`
	Insight  uint32 `json:"insight"`
	Mediocre uint32 `json:"mediocre"`
}

type News struct {
	ID             string         `json:"id"`
	Category       string         `json:"category"`
	Title          string         `json:"title"`
	Source         string         `json:"source"`
	URL            string         `json:"url"`
	Summary        string         `json:"summary"`
	PublishedAt    time.Time      `json:"published_at"`
	FeedbackScores FeedbackScores `json:"feedback_scores"`
	CommentIDs     []string       `json:"comment_ids"`
}
