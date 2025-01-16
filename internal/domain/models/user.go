package models

type User struct {
	ID                  string   `json:"id"`
	CommentIDs          []string `json:"comment_ids"`
	Gender              string   `json:"gender"`
	AgeGroup            uint32   `json:"age_group"`
	Occupation          string   `json:"occupation"`
	PoliticalView       string   `json:"political_view"`
	OpinionTone         string   `json:"opinion_tone"`
	SpeechStyle         string   `json:"speech_style"`
	CommentLength       uint32   `json:"comment_length"`
	BackgroundKnowledge string   `json:"background_knowledge"`
	Emotion             string   `json:"emotion"`
}
