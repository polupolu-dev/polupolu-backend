package validater

import (
	"errors"
	"net/url"
)

// URL 文字列のバリデーション
func IsValidURL(rawURL string) error {
	// URL のパース
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	// スキームとホストが含まれているか確認
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return errors.New("scheme and host not included")
	}

	return nil
}
