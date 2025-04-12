package misskey

// Client はMisskey APIとやり取りするためのクライアント
type Client struct {
	ApiToken string
	ApiUrl   string
}

// NewClient はMisskeyクライアントを初期化する
func NewClient(apiToken, apiUrl string) *Client {
	return &Client{
		ApiToken: apiToken,
		ApiUrl:   apiUrl,
	}
}
