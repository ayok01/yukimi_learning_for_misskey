package usecase

import (
	"log"
	"math/rand"
	"time"

	"github.com/ayok01/yukimi_learning_for_misskey/misskey"
)

type NoteUsecase struct {
	Client *misskey.Client
}

func NewNoteUsecase(client *misskey.Client) *NoteUsecase {
	return &NoteUsecase{Client: client}
}

// GetRandomNote はタイムラインからランダムなノートを取得します
func (u *NoteUsecase) GetRandomNote(request misskey.TimelineRequest, timelineType string) (*misskey.Note, error) {
	notes, err := u.Client.GetTimeline(request, timelineType)
	if err != nil {
		return nil, err
	}

	if len(notes) == 0 {
		log.Println("No notes available.")
		return nil, nil
	}

	// ランダムなノートを取得
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNote := notes[r.Intn(len(notes))]

	//　ノートのテキストが空でないことを確認
	if randomNote.Text == "" {
		log.Println("Random note text is empty.")
		return nil, nil
	}
	// noteの公開設定が"public"でない場合はnilを返す
	if randomNote.Visibility != "public" {
		log.Println("Random note visibility is not public.")
		return nil, nil
	}
	return &randomNote, nil
}
