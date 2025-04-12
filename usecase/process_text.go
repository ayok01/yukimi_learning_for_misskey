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

	// テキストの中にURLが含まれている場合はnilを返す
	if yukimi_text.ContainsURL(processedText) {
		log.Println("Processed text contains URL.")
		return "", nil
	}

	// テキストの中に:emoji_name:の形式の絵文字が含まれている場合はnilを返す
	if yukimi_text.ContainsEmoji(processedText) {
		log.Println("Processed text contains emoji.")
		return "", nil
	}

	// テキストの中に@user_nameの形式のユーザー名が含まれている場合はnilを返す
	if yukimi_text.ContainsUserName(processedText) {
		log.Println("Processed text contains user name.")
		return "", nil
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
