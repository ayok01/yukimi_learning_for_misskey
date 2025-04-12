package misskey

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type CreateFollowRequest struct {
	UserID      string `json:"userId"`
	I           string `json:"i"`
	WithReplies bool   `json:"withReplies"` // フォロワーのリプライを取得するかどうか
}

func (c *Client) CreateFollow(request *CreateFollowRequest) error {
	url := "https://" + c.ApiUrl + "/api/following/create"

	// リクエストボディを作成
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	// HTTP POSTリクエストを送信
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// ステータスコードを確認
	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to fetch follow: " + resp.Status)
	}

	// レスポンスをパース
	var notes []Note
	if err := json.NewDecoder(resp.Body).Decode(&notes); err != nil {
		return err
	}

	return nil
}
