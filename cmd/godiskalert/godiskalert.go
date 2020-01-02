package main

import (
	"flag"
	"fmt"
	"github.com/take64/godiskalert/internal/apps/godiskalert"
	"os"
	"time"
)

func main() {

	// 引数
	var slack = flag.String("slack", "", "slack webhook URL")
	flag.Parse()
	if *slack == "" {
		fmt.Println("-slack 引数がありません")
		os.Exit(1)
	}

	// ディレクトリ取得
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// ディスク情報を取得
	diskInfo := godiskalert.Info(wd)
	formattedDiskInfo := diskInfo.Format()

	// slackペイロードの生成
	var fields []godiskalert.Fields
	fields = append(fields, godiskalert.Fields{
		Title: "IPアドレス",
		Value: godiskalert.IpAddress(),
		Short: true,
	})
	fields = append(fields, godiskalert.Fields{
		Title: "ディスク残り容量",
		Value: fmt.Sprintf("%s(%s)", formattedDiskInfo.Free, formattedDiskInfo.FreePercent),
		Short: true,
	})

	var attachments []godiskalert.Attachments
	attachments = append(attachments, godiskalert.Attachments{
		Fallback:   "",
		Color:      "FF9933",
		Pretext:    "",
		AuthorName: "",
		AuthorLink: "",
		AuthorIcon: "",
		Title:      "容量確認",
		TitleLink:  "",
		Text:       "<!here>",
		Fields:     fields,
		ImageUrl:   "",
		ThumbUrl:   "",
		Footer:     "",
		FooterIcon: "",
		Ts:         time.Now().Unix(),
	})

	payload := godiskalert.Payload{
		Attachments:attachments,
	}

	// slackにポスト
	err = godiskalert.WebhookPost(*slack, payload)
}
