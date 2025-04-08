package main

import (
    "log"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func startBot() {
    bot, err := tgbotapi.NewBotAPI("YOUR_BOT_TOKEN")
    if err != nil {
        log.Fatal(err)
    }

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil {
            continue
        }

        btn := tgbotapi.NewKeyboardButton("🛒 Открыть магазин")
        btn.WebApp = &tgbotapi.WebAppInfo{URL: "https://your-frontend-url.com"}

        keyboard := tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{btn})
        keyboard.ResizeKeyboard = true

        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Нажми кнопку для входа в магазин:")
        msg.ReplyMarkup = keyboard
        bot.Send(msg)
    }
}
