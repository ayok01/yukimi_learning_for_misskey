package misskey

import (
	"encoding/json"
	"log"

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
	// WebSocket接続の実装
	url := "wss://" + c.ApiUrl + "/streaming?i=" + c.ApiToken

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v", err)
		return err
	}
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	// ユーザーチャンネルに接続
	channelID := "user-follow"
	connectMessage := map[string]interface{}{
		"type": "connect",
		"body": map[string]interface{}{
			"channel": "main",
			"id":      channelID,
		},
	}

	if err := conn.WriteJSON(connectMessage); err != nil {
		log.Fatalf("チャンネル接続メッセージの送信に失敗しました: %v", err)
	}

	log.Println("フォローイベントの待機を開始します...")

	// メッセージを受信
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("メッセージの受信中にエラーが発生しました: %v", err)
			break
		}

		// メッセージを解析
		var response map[string]interface{}
		if err := json.Unmarshal(message, &response); err != nil {
			log.Printf("メッセージの解析に失敗しました: %v", err)
			continue
		}

		// フォローイベントをチェック
		if response["type"] == "channel" {
			body, ok := response["body"].(map[string]interface{})
			if !ok {
				log.Printf("bodyの解析に失敗しました: %v", response)
				continue
			}

			if body["id"] == channelID && body["type"] == "notification" {
				innerBody, ok := body["body"].(map[string]interface{})
				if !ok {
					log.Printf("innerBodyの解析に失敗しました: %v", body)
					continue
				}

				if innerBody["type"] == "follow" {
					// userIdの取得
					userId, ok := innerBody["userId"].(string)
					if !ok {
						log.Printf("フォローイベントのuserIdが見つかりません: %v", innerBody)
						continue
					}

					// ユーザー情報の取得
					user, ok := innerBody["user"].(map[string]interface{})
					if !ok {
						log.Printf("フォローイベントのuser情報が見つかりません: %v", innerBody)
						continue
					}

					username, _ := user["username"].(string)
					host, _ := user["host"].(string)

					log.Printf("フォローイベントを受信しました: ユーザーID: %v, ユーザー名: %v, ホスト: %v", userId, username, host)

					// フォロー処理
					c.CreateFollow(&CreateFollowRequest{
						UserID:      userId,
						I:           c.ApiToken,
						WithReplies: false,
					})
				}
			}
		}
	}

	return nil
}
