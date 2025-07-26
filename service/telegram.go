package service

import (
	"fmt"
	"log"
	"strconv"
	"trust-credit-back/environment"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	Bot       *tgbotapi.BotAPI
	ChannelID int64
)

func InitTelegramBot(token string) error {
	var err error
	Bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	ch := environment.GetVariable("TELEGRAM_CHANNEL_ID")
	if ch != "" {
		if id, err := strconv.ParseInt(ch, 10, 64); err == nil {
			ChannelID = id
		} else {
			log.Println("Invalid TELEGRAM_CHANNEL_ID:", err)
		}
	}
	return nil
}

func SendTelegramCode(chatID int64, code string) error {
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Ваш код авторизации: %s", code))
	_, err := Bot.Send(msg)
	return err
}

func SendCodeToUserOrChannel(chatID int64, code string) error {
	if chatID > 0 {
		msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Ваш код: %s", code))
		_, err := Bot.Send(msg)
		return err
	}

	return SendCodeToChannel(code)
}

func SendCodeToChannel(code string) error {
	if ChannelID == 0 {
		return fmt.Errorf("channel ID not initialized")
	}
	msg := tgbotapi.NewMessage(ChannelID, fmt.Sprintf("Ваш код: %s", code))
	_, err := Bot.Send(msg)
	return err
}
