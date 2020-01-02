package godiskalert

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// slackペイロードのfield
type Fields struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

// slackペイロードのattachments
type Attachments struct {
	Fallback   string   `json:"fallback"`
	Color      string   `json:"color"`
	Pretext    string   `json:"pretext"`
	AuthorName string   `json:"author_name"`
	AuthorLink string   `json:"author_link"`
	AuthorIcon string   `json:"author_icon"`
	Title      string   `json:"title"`
	TitleLink  string   `json:"title_link"`
	Text       string   `json:"text"`
	Fields     []Fields `json:"fields"`
	ImageUrl   string   `json:"image_url"`
	ThumbUrl   string   `json:"thumb_url"`
	Footer     string   `json:"footer"`
	FooterIcon string   `json:"footer_icon"`
	Ts         int64    `json:"ts"`
}

// slackペイロード
type Payload struct {
	Attachments []Attachments `json:"attachments"`
}

// slack webhookにPOSする
func WebhookPost(url string, attachments Payload) error {
	jsonBytes, err := json.Marshal(attachments)
	//fmt.Println(string(jsonBytes))
	req, err := http.NewRequest(
		"POST",
		url,
		strings.NewReader(string(jsonBytes)),
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")

	// 接続
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	return err
}
