package main

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/yuki0920/weather_line_bot_sample/weather"
)

func main() {

	// LINE Botクライアント生成する
	// BOT にはチャネルシークレットとチャネルトークンを環境変数から読み込み引数に渡す
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)

	// エラーに値があればログに出力し終了する
	if err != nil {
		log.Fatal(err)
	}

	// weatherパッケージパッケージから天気情報の文字列をを取得する
	result, err := weather.GetWeather()

	if err != nil {
		log.Fatal(err)
	}

	message := linebot.NewTextMessage(result)

	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}