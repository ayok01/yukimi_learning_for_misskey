package misskey

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type WebSocketResponse struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func (c *Client) WebSocketConnect() error {
	const retryInterval = 5 * time.Second

	for {
		// WebSocket接続を試みる
		conn, err := c.connectWebSocket()
		if err != nil {
			log.Printf("WebSocket接続に失敗しました: %v", err)
			log.Printf("再接続を %v 後に試みます...", retryInterval)
			time.Sleep(retryInterval)
			continue
		}

		// メッセージの受信を開始
		if err := c.handleWebSocketMessages(conn); err != nil {
			log.Printf("WebSocketエラー: %v", err)
			log.Println("再接続を試みます...")
			time.Sleep(retryInterval)
			continue
		}
	}
}

func (c *Client) connectWebSocket() (*websocket.Conn, error) {
	url := "wss://" + c.ApiUrl + "/streaming?i=" + c.ApiToken
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	log.Println("WebSocket接続に成功しました")
	return conn, nil
}

func (c *Client) handleWebSocketMessages(conn *websocket.Conn) error {
	defer conn.Close()

	channelID := "user-follow"
	connectMessage := map[string]interface{}{
		"type": "connect",
		"body": map[string]interface{}{
			"channel": "main",
			"id":      channelID,
		},
	}

	if err := conn.WriteJSON(connectMessage); err != nil {
		return err
	}

	log.Println("フォローイベントの待機を開始します...")

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			return err // エラーが発生した場合、再接続をトリガー
		}

		// メッセージの処理
		var response map[string]interface{}
		if err := json.Unmarshal(message, &response); err != nil {
			log.Printf("メッセージの解析に失敗しました: %v", err)
			continue
		}

		// フォローイベントの処理
		c.processFollowEvent(response, channelID)
	}
}

func (c *Client) processFollowEvent(response map[string]interface{}, channelID string) {
	if response["type"] == "channel" {
		body, ok := response["body"].(map[string]interface{})
		if !ok {
			log.Printf("bodyの解析に失敗しました: %v", response)
			return
		}

		if body["id"] == channelID && body["type"] == "notification" {
			innerBody, ok := body["body"].(map[string]interface{})
			if !ok {
				log.Printf("innerBodyの解析に失敗しました: %v", body)
				return
			}

			if innerBody["type"] == "follow" {
				userId, _ := innerBody["userId"].(string)
				user, _ := innerBody["user"].(map[string]interface{})
				username, _ := user["username"].(string)
				host, _ := user["host"].(string)

				log.Printf("フォローイベントを受信しました: ユーザーID: %v, ユーザー名: %v, ホスト: %v", userId, username, host)

				c.CreateFollow(&CreateFollowRequest{
					UserID:      userId,
					I:           c.ApiToken,
					WithReplies: false,
				})
			}
		}
	}
}
