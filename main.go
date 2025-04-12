package main

import (
	"log"
	"os"

	"github.com/ayok01/yukimi_learning_for_misskey/batch"
	"github.com/ayok01/yukimi_learning_for_misskey/config"
	"github.com/ayok01/yukimi_learning_for_misskey/misskey"
	"github.com/ayok01/yukimi_learning_for_misskey/yukimi_text"
	"github.com/bluele/mecab-golang"
	"github.com/joho/godotenv"
)

func main() {
	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数から設定を読み込む
	apiToken := os.Getenv("MISSKEY_API_TOKEN")
	apiUrl := os.Getenv("MISSKEY_API_URL")

	if apiToken == "" || apiUrl == "" {
		log.Fatal("Environment variables MISSKEY_API_TOKEN or MISSKEY_API_URL are not set")
	}

	cfg := config.Config{
		ApiToken: apiToken,
		ApiUrl:   apiUrl,
	}

	// Misskeyクライアントを初期化
	client := misskey.NewClient(cfg.ApiToken, cfg.ApiUrl)

	// WebSocket接続を並行して実行
	go func() {
		err = client.WebSocketConnect()
		if err != nil {
			log.Fatalf("Error connecting to WebSocket: %v", err)
		}
	}()

	// MeCabを使ったテキストプロセッサを初期化
	m, err := mecab.New("-Owakati")
	if err != nil {
		log.Fatalf("Error initializing MeCab: %v", err)
	}
	defer m.Destroy()

	textProcessor := yukimi_text.NewYukimiTextProcessor(m)

	// バッチ処理を開始
	batch.ProcessTimeline(client, textProcessor)
}
