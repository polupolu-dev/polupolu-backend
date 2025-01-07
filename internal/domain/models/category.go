package models

type Category struct {
	// omitempty はデコード時に値が空だと省略してくれる
	Name    string   `json:"name"`
	NewsIDs []string `json:"news_ids,omitempty"`
}
