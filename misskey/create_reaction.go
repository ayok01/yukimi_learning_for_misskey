package misskey

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type CreateReactionRequest struct {
	NoteID   string `json:"noteId"`
	Reaction string `json:"reaction"`
	I        string `json:"i"`
}

func (c *Client) CreateReaction(request CreateReactionRequest) error {
	url := "https://" + c.ApiUrl + "/api/notes/reactions/create"

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
		return errors.New("failed to fetch reaction: " + resp.Status)
	}

	return nil
}
