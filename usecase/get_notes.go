package usecase

import (
	"log"
	"math/rand"

	"github.com/ayok01/yukimi_learning_for_misskey/misskey"
	"github.com/ayok01/yukimi_learning_for_misskey/model"
)

type NoteUsecase struct {
	Client *misskey.Client
}

func NewNoteUsecase(client *misskey.Client) *NoteUsecase {
	return &NoteUsecase{Client: client}
}

// GetRandomNote はタイムラインからランダムなノートを取得します
func (u *NoteUsecase) GetRandomNote(request misskey.TimelineRequest, timelineType string, me *model.User) (*misskey.Note, error) {
	notes, err := u.Client.GetTimeline(request, timelineType)
	if err != nil {
		return nil, err
	}

	if len(notes) == 0 {
		log.Println("No notes available.")
		return nil, nil
	}

	log.Println("Number of notes:", len(notes))

	// ランダムなノートを取得する処理を最大3回試行
	var randomNote *misskey.Note
	for i := 0; i < 3; i++ {
		randomIndex := rand.Intn(len(notes)) // グローバルな乱数生成器を使用
		candidateNote := notes[randomIndex]

		// ノートのユーザーが自分でないことを確認
		log.Println("Random note userID:", candidateNote.User.UserID, me.UserID)
		if candidateNote.User.UserID == me.UserID {
			log.Println("Random note is from the user itself.")
			continue
		}

		// 前回のノートと同じでないことを確認
		if candidateNote.MyReaction != "" {
			log.Println("Random note is already reacted.")
			continue
		}

		// ノートのテキストが空でないことを確認
		if candidateNote.Text == "" {
			log.Println("Random note text is empty.")
			continue
		}

		// noteの公開設定が"public"でない場合はスキップ
		if candidateNote.Visibility != "public" && candidateNote.LocalOnly {
			log.Println("Random note visibility is not public.", candidateNote.Visibility, candidateNote.Text, candidateNote.ID)
			continue
		}

		// 条件を満たすノートが見つかった場合
		randomNote = &candidateNote
		break
	}

	if randomNote == nil {
		log.Println("No suitable random note found after 3 attempts.")
		return nil, nil
	}

	return randomNote, nil
}
