package misskey

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type CreateNoteRequest struct {
	Text       string `json:"text"`
	Visibility string `json:"visibility"`
	// 追加のフィールドが必要な場合はここに追加
}

func NewCreateNoteRequest(text, visibility string) *CreateNoteRequest {
	return &CreateNoteRequest{
		Text:       text,
		Visibility: visibility,
	}
}

func (c *Client) CreateNote(request *CreateNoteRequest) error {
	url := "https://" + c.ApiUrl + "/api/notes/create"

	// リクエストボディを作成
	body, err := json.Marshal(request)
	if err != nil {
		log.Printf("Error marshaling request body: %v", err)
		return err
	}

	// HTTP POSTリクエストを送信
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: received status code %d", resp.StatusCode)
		return errors.New("failed to create note")
	}

	return nil
}
