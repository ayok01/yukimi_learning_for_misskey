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

	// 自分のユーザー情報を取得
	user, err := client.GetMe(misskey.GetMeRequest{I: client.ApiToken})
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return
	}

	// 10分ごとに処理を実行するループ
	for {
		log.Println("Starting batch process...")

		// タイムラインリクエストを作成
		request := misskey.TimelineRequest{
			WithFiles:    false,
			WithRenotes:  false,
			WithReplies:  false,
			Limit:        10,
			AllowPartial: true,
		}

		// ランダムなノートを取得
		randomNote, err := noteUsecase.GetRandomNote(request, "home", user)
		if err != nil {
			log.Printf("Error getting random note: %v", err)
		} else if randomNote == nil || randomNote.Text == "" {
			log.Println("No valid random note available or text is empty.")
		} else {
			log.Printf("Random note ID: %s, Text: %s", randomNote.ID, randomNote.Text)
			// ノートのテキストを加工
			processedText, err := textProcessorUsecase.ProcessNoteText(randomNote.Text)
			if err != nil {
				log.Printf("Error processing text: %v", err)
			} else if processedText == "" {
				log.Println("Processed text is empty.")
			} else {
				// ノートにリアクションを追加
				createReactionRequest := misskey.CreateReactionRequest{
					NoteID:   randomNote.ID,
					Reaction: "❤️",
					I:        client.ApiToken,
				}
				log.Println("Creating reaction...", createReactionRequest.NoteID)
				err = client.CreateReaction(createReactionRequest)
				if err != nil {
					log.Printf("Error creating reaction: %v", err)
				} else {
					log.Println("Reaction created successfully.")
				}
				// ノートを投稿
				note := misskey.CreateNoteRequest{
					Text:       processedText,
					Visibility: "public",
					I:          client.ApiToken,
				}
				err = client.CreateNote(note)
				if err != nil {
					log.Printf("Error creating note: %v", err)
				} else {
					log.Println("Note created successfully.")
				}
				log.Println("Creating note with processed text...", processedText)
			}
		}

		// 10分待機
		log.Println("Batch process completed. Waiting for 10 minutes...")
		time.Sleep(10 * time.Minute)
	}
}
