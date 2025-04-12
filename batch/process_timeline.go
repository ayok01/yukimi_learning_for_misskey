package batch

import (
	"fmt"
	"log"
	"time"

	"github.com/ayok01/yukimi_learning_for_misskey/misskey"
	"github.com/ayok01/yukimi_learning_for_misskey/usecase"
	"github.com/ayok01/yukimi_learning_for_misskey/yukimi_text"
)

// ProcessTimeline はタイムラインを取得してノートを加工するバッチ処理
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
		randomNote, err := noteUsecase.GetRandomNote(request, "local")
		if err != nil {
			log.Printf("Error getting random note: %v", err)
			continue
		}
		if randomNote == nil {
			log.Println("No random note available.")
			continue
		}

		// noteが空の場合はスキップ
		if randomNote.Text == "" {
			log.Println("Random note text is empty.")
			continue
		}

		// ノートのテキストを加工
		processedText, err := textProcessorUsecase.ProcessNoteText(randomNote.Text)
		if err != nil {
			log.Printf("Error processing text: %v", err)
			continue
		}

		// 加工後のテキストを出力
		fmt.Printf("Processed Text: %s\n", processedText)

		// 10分待機
		log.Println("Batch process completed. Waiting for 10 minutes...")
		time.Sleep(10 * time.Minute)
	}
}
