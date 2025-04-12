package yukimi_text

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/bluele/mecab-golang"
)

const BOSEOS = "BOS/EOS"

// YukimiTextProcessor はテキスト処理を行う構造体
type YukimiTextProcessor struct {
	MeCab *mecab.MeCab
}

// NewYukimiTextProcessor は新しい YukimiTextProcessor を初期化する
func NewYukimiTextProcessor(m *mecab.MeCab) *YukimiTextProcessor {
	return &YukimiTextProcessor{
		MeCab: m,
	}
}

// ChangeYukimiText はテキストを加工するメソッド
func (p *YukimiTextProcessor) ChangeYukimiText(text string) (string, error) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	var analyzedTweets []string

	tg, err := p.MeCab.NewTagger()
	if err != nil {
		return "", err
	}
	defer tg.Destroy()

	lt, err := p.MeCab.NewLattice(text)
	if err != nil {
		return "", err
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] != BOSEOS {
			fmt.Printf("%s %s\n", node.Surface(), node.Feature())
		}
		if node.Surface() != "" {
			analyzedTweets = append(analyzedTweets, node.Surface())
		}

		if len(features) > 0 {
			partOfSpeech := features[0]
			// 副詞または助詞の場合、25%の確率で三点リーダを付ける
			if partOfSpeech == "副詞" || partOfSpeech == "助詞" {
				log.Println("three dots")
				if rng.Intn(4) == 0 {
					log.Println("…")
					for i := 0; i < rand.Intn(4)+1; i++ {
						analyzedTweets = append(analyzedTweets, "…")
					}
				}
			}
		}

		if node.Next() != nil {
			break
		}
	}

	// 12.5%の確率で文末に「ふふ」を追加
	if rand.Intn(8) == 0 {
		analyzedTweets = append(analyzedTweets, strings.Repeat("…", rand.Intn(4)+1)+"ふふ"+strings.Repeat("…", rand.Intn(4)+1))
	}

	result := strings.Join(analyzedTweets, "")
	log.Println(result)
	return result, nil
}
