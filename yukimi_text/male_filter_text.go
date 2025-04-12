package yukimi_text

import (
	"bufio"
	"os"
	"strings"
)

// LoadNGWords はNGワードリストをファイルから読み込みます
func LoadNGWords(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ngWords []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" && !strings.HasPrefix(word, "//") { // コメント行を無視
			ngWords = append(ngWords, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ngWords, nil
}

// ContainsNGWord はテキストにNGワードが含まれているかをチェックします
func ContainsNGWord(text string, ngWords []string) bool {
	for _, word := range ngWords {
		if strings.Contains(text, word) {
			return true
		}
	}
	return false
}

// ContainsURL はテキストにURLが含まれているかをチェックします
func ContainsURL(text string) bool {
	// URLの正規表現パターン
	urlPattern := `(?i)\b(?:https?://|www\.)\S+\b`
	return strings.Contains(text, urlPattern)
}

// テキストの中に:emoji_name:の形式の絵文字が含まれているかをチェックします
func ContainsEmoji(text string) bool {
	// :emoji_name:の正規表現パターン
	emojiPattern := `:[a-zA-Z0-9_]+:`
	return strings.Contains(text, emojiPattern)
}

// @user_nameの形式のユーザー名が含まれているかをチェックします
func ContainsUserName(text string) bool {
	// @user_nameの正規表現パターン
	userNamePattern := `@[a-zA-Z0-9_]+`
	return strings.Contains(text, userNamePattern)
}
