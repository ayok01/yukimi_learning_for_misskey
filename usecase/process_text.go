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

	// NGワードを含む場合はnilを返す
	ngWordFilePath := "./data/ngword.txt"
	ngWords, err := yukimi_text.LoadNGWords(ngWordFilePath)
	if err != nil {
		log.Printf("Error loading NG words: %v", err)
		return "", err
	}
	// NGワードを含むかどうかをチェック
	judge := yukimi_text.ContainsNGWord(processedText, ngWords)
	if judge {
		log.Println("Processed text contains NG words.")
		return "", nil
	}

	return processedText, nil
}
