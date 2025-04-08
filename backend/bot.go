import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

bot, _ := tgbotapi.NewBotAPI("YOUR_BOT_TOKEN")

msg := tgbotapi.NewMessage(chatID, "Открой мини-приложение:")
msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{
    Keyboard: [][]tgbotapi.KeyboardButton{
        {
            {
                Text: "Открыть магазин",
                WebApp: &tgbotapi.WebAppInfo{
                    URL: "https://your-frontend-url.com",
                },
            },
        },
    },
    ResizeKeyboard: true,
}

bot.Send(msg)

