package misskey

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ayok01/yukimi_learning_for_misskey/model"
)

type GetUserRequest struct {
	UserID string `json:"id"`
	I      string `json:"i"`
}

func (c *Client) GetUser(request GetUserRequest) (*model.User, error) {
	url := "https://" + c.ApiUrl + "/api/users/show"

	// リクエストボディを作成
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
		return nil, errors.New("failed to fetch user: " + resp.Status)
	}

	// レスポンスをパース
	var user model.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

type GetMeRequest struct {
	I string `json:"i"`
}

func (c *Client) GetMe(request GetMeRequest) (*model.User, error) {
	url := "https://" + c.ApiUrl + "/api/i"

	// リクエストボディを作成
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
		return nil, errors.New("failed to fetch user: " + resp.Status)
	}

	// レスポンスをパース
	var user model.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
