package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
)

var infoKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		//tgbotapi.NewInlineKeyboardButtonSwitch("What can you do?", "What can you do?"),
		tgbotapi.NewInlineKeyboardButtonData("What can you do?", "Info"),
		tgbotapi.NewInlineKeyboardButtonData("Repayment", "Repayment"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Cashback", "Cashback"),
		tgbotapi.NewInlineKeyboardButtonData("Fee", "Fee"),
		tgbotapi.NewInlineKeyboardButtonData("Replenishment", "Replenishment"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Credit", "Credit"),
	),
)

var canDoKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Example1", "Example1"),
		tgbotapi.NewInlineKeyboardButtonData("Example2", "Example2"),
		tgbotapi.NewInlineKeyboardButtonData("Example3", "Example3"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1475925546:AAEP3tiMJ9ORDcNECfIX0v7wihG07dso3rY")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		//UserName := update.Message.From.UserName
		//UserID := update.Message.From.ID
		//ChatID := update.Message.Chat.ID
		//Text := update.Message.Text

		//log.Printf("[%s] %d %d %s", UserName, UserID, ChatID, Text)

		if update.CallbackQuery != nil {
			//fmt.Print(updates)

			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID,update.CallbackQuery.Data))

			queryMSG := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID,update.CallbackQuery.Data)

			bot.Send(queryMSG)
		}

		if update.ChosenInlineResult != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.ChosenInlineResult. {
			case "Info":
				//video := tgbotapi.NewVideoShare(ChatID, "")
				msg.Text = "I can a lot))"
				msg.ReplyToMessageID = update.Message.MessageID
				msg.ReplyMarkup = canDoKeyboard
			}

			bot.Send(msg)
		}

		if update.InlineQuery != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.InlineQuery. {
			case "Info":
				//video := tgbotapi.NewVideoShare(ChatID, "")
				msg.Text = "I can a lot))"
				msg.ReplyToMessageID = update.Message.MessageID
				msg.ReplyMarkup = canDoKeyboard
			}

			bot.Send(msg)
		}
		if update.ChannelPost != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.ChannelPost.Text {
			case "Info":
				//video := tgbotapi.NewVideoShare(ChatID, "")
				msg.Text = "I can a lot))"
				msg.ReplyToMessageID = update.Message.MessageID
				msg.ReplyMarkup = canDoKeyboard
			}

			bot.Send(msg)
		}

		if update.Message != nil {
			//msg := tgbotapi.NewMessage(ChatID, Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)


			switch update.Message.Command() {
			case "help":
				msg.Text = "type /hi or /status."
			case "hi":
				msg.Text = "Hi :)"
			case "status":
				msg.Text = "I'm ok."
			default:
				msg.Text = "I don't know that command"
			}

			switch update.Message.Text {
			case "/start":
				msg.Text = "Hello"
				msg.ReplyToMessageID = update.Message.MessageID
				msg.ReplyMarkup = infoKeyboard
			case "Info":
				//video := tgbotapi.NewVideoShare(ChatID, "")
				msg.Text = "I can a lot))"
				msg.ReplyToMessageID = update.Message.MessageID
				msg.ReplyMarkup = canDoKeyboard
			default:
				msg.Text = ""
			}

			bot.Send(msg)
		}
	}
}
