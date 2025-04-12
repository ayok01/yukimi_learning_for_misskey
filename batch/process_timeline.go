package batch

import (
	"log"
	"time"

	"github.com/ayok01/yukimi_learning_for_misskey/misskey"
	"github.com/ayok01/yukimi_learning_for_misskey/usecase"
	"github.com/ayok01/yukimi_learning_for_misskey/yukimi_text"
)

// ProcessTimeline はタイムラインを取得してノートを加工して投稿するバッチ処理を行う関数
func ProcessTimeline(client *misskey.Client, textProcessor *yukimi_text.YukimiTextProcessor) {
	// ノート取得ユースケースを初期化
	noteUsecase := usecase.NewNoteUsecase(client)

	// テキスト加工ユースケースを初期化
	textProcessorUsecase := usecase.NewTextProcessorUsecase(textProcessor)

	// 10分ごとに処理を実行するループ
	for {
		log.Println("Starting batch process...")

		// タイムラインリクエストを作成
		request := misskey.TimelineRequest{
			WithFiles:    true,
			WithRenotes:  true,
			WithReplies:  true,
			Limit:        10,
			AllowPartial: true,
		}

		// ランダムなノートを取得
		randomNote, err := noteUsecase.GetRandomNote(request, "home")
		if err != nil {
			log.Printf("Error getting random note: %v", err)
		} else if randomNote == nil || randomNote.Text == "" {
			log.Println("No valid random note available or text is empty.")
		} else {
			// ノートのテキストを加工
			processedText, err := textProcessorUsecase.ProcessNoteText(randomNote.Text)
			if err != nil {
				log.Printf("Error processing text: %v", err)
			} else if processedText == "" {
				log.Println("Processed text is empty.")
			} else {
				// ノートを投稿
				note := misskey.CreateNoteRequest{
					Text:       processedText,
					Visibility: "public",
				}
				err = client.CreateNote(note)
				if err != nil {
					log.Printf("Error creating note: %v", err)
				} else {
					log.Println("Note created successfully.")
				}
			}
		}

		// 10分待機
		log.Println("Batch process completed. Waiting for 10 minutes...")
		time.Sleep(10 * time.Minute)
	}
}
