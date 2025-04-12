package misskey

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// TimelineRequest はタイムライン取得のリクエスト構造体
type TimelineRequest struct {
	WithFiles    bool   `json:"withFiles"`
	WithRenotes  bool   `json:"withRenotes"`
	WithReplies  bool   `json:"withReplies"`
	Limit        int    `json:"limit"`
	AllowPartial bool   `json:"allowPartial"`
	I            string `json:"i"`
}

// Note はMisskeyのノート構造体
type Note struct {
	ID         string `json:"id"`
	Text       string `json:"text"`
	Visibility string `json:"visibility"`
	LocalOnly  bool   `json:"localOnly"`
}

// GetTimeline はタイムラインを取得する
func (c *Client) GetTimeline(request TimelineRequest, timelineType string) ([]Note, error) {
	url := "https://" + c.ApiUrl + "/api/notes/timeline"
	if timelineType == "local" {
		url = "https://" + c.ApiUrl + "/api/notes/local-timeline"
	}

	// リクエストボディを作成
	request.I = c.ApiToken
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// HTTP POSTリクエストを送信
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// ステータスコードを確認
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch timeline: " + resp.Status)
	}

	// レスポンスをパース
	var notes []Note
	if err := json.NewDecoder(resp.Body).Decode(&notes); err != nil {
		return nil, err
	}

	return notes, nil
}
