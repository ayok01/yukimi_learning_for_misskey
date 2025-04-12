package usecase

import (
	"log"

	"github.com/ayok01/yukimi_learning_for_misskey/yukimi_text"
)

type TextProcessorUsecase struct {
	Processor *yukimi_text.YukimiTextProcessor
}

func NewTextProcessorUsecase(processor *yukimi_text.YukimiTextProcessor) *TextProcessorUsecase {
	return &TextProcessorUsecase{Processor: processor}
}

// ProcessNoteText はノートのテキストを加工します
func (u *TextProcessorUsecase) ProcessNoteText(text string) (string, error) {
	if text == "" {
		log.Println("Note text is empty.")
		return "", nil
	}

	processedText, err := u.Processor.ChangeYukimiText(text)
	if err != nil {
		return "", err
	}
	return processedText, nil
}
