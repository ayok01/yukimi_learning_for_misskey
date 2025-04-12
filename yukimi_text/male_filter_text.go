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
