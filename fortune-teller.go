package main

import (
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const TOKEN = "7732419880:AAFauADWEye0Eq4h5ScXxx-KM1Dy_3QAs8M"

var chatid int64
var bot *tgbotapi.BotAPI

var fortuneTellerNames = [3]string{"ГРЕПЛЕР228", "ромчиксамбист", "фирдавскриптан"}

var answer = []string{
	"пока что нет",
	"Ты уверен, что хочешь узнать правду? ",
	"мог бы и не спрашивать ",
	"Бывают ли вторые шансы? ",
	"Ты готов взять на себя ответственность? ",
	"очевидно-ДА",
	"Могут ли прошлые ошибки охватить тебя? ",
	"Человек может измениться? ",
	"Стоит ли рисковать ради любви? ",
	"Ты готов противостоять своим страхам? ",
	"Может ли судьба быть предрешенной? ",
	"Ты боишься смотреть в глаза правде? ",
	"Ты готов отпустить прошлое? (Зеленая миля)",
	"Может ли человек убежать от своей судьбы? (Темная башня: Стрелок)",
}

func connectwithTelegram() {
	var err error
	if bot, err = tgbotapi.NewBotAPI(TOKEN); err != nil {
		panic("нет подключения к телеграмм")
	}
}
func sendMessage(msg string) {
	msgConfig := tgbotapi.NewMessage(chatid, msg)
	bot.Send(msgConfig)
}
func isMessageForFortuneTeller(update *tgbotapi.Update) bool {
	if update.Message == nil || update.Message.Text == "" {
		return false
	}
	msgInLowerCase := strings.ToLower(update.Message.Text)
	for _, name := range fortuneTellerNames {
		if strings.Contains(msgInLowerCase, name) {
			return true
		}
	}
	return false
}
func getFortuneTellersAnswer() string {
	index := rand.Intn(len(answer))
	return answer[index]
}
func sendAnswer(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(chatid, getFortuneTellersAnswer())
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}
func main() {
	connectwithTelegram()

	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatid = update.Message.Chat.ID
			sendMessage("задай свой вопрос назвав меня по имени. Ответом должн быть либо \"да\" либо \"нет\". Например  \"Грпеплер228, Тихон о" +
				"тлижет Хроль в этом семестре?\"")
		}

		if isMessageForFortuneTeller(&update) {
			sendAnswer(&update)
		}
	}
}
