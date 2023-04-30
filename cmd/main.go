package main

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/KarlvenK/tgBot/config"
	"github.com/KarlvenK/tgBot/pkg/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	defer log.Flush()
	bot, err := tgbotapi.NewBotAPI(config.GetConfig().Token)
	if err != nil {
		log.Error(fmt.Errorf("fail to create bot instance, %v", err))
	}
	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		log.Info("get a msg: ", msg.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(msg); err != nil {
			log.Error(errors.Errorf("fail to send msg, %v", err))
		}
	}
}
