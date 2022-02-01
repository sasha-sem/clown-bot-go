package handlers

import (
	"clownbot/internal/services/wordle"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func (hs *HandlerService) WordleHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	content := ""
	switch i.ApplicationCommandData().Options[0].Name {
	case "new":
		err := hs.wordleService.New()
		content = "🤡 Я загадал слово из 5 букв, попробуйте отгадать 🤡"
		if err != nil {
			content = "🥲У меня не получилось загадать слово, попробуй снова использовать команду!🥲"
		}
	case "help":
		content = ">>> Угадай загаданное слово из шести букв.Для того, чтобы я загадал слово введи команду /wordle new\n" +
			"Чтобы попробовать отгадать слово введи команду /wordle guess *слово из шести букв*\n" +
			"В ответ бот выдаст тебе последовательность эмоджи.\n" +
			"Если буква в твоем слове стоит на том же месте, что и в загаданном, то я отмечу ее положение зеленым :green_square:," +
			" если буква есть в загаданом слове, но стоит на неверном месте, то я отмечу ее положение оранжевым :orange_square:," +
			" если же буквы нет в загаданом слове, то я отмечу ее положение белым :white_large_square:.\n" +
			"🤡 Уверен, у тебя не получится отгадать, загаданное мною слово! Удачи! 🤡"

	case "guess":
		word := strings.TrimSpace(i.ApplicationCommandData().Options[0].Options[0].StringValue())
		word = strings.ToLower(word)

		if len([]rune(word)) != 5 {
			content = "Слово должно быть из 5 букв 🤡"
			break
		}

		answer, err := hs.wordleService.Guess(word)
		if err == wordle.WordWasNotSet {
			content = "🤡 Слово не было загадано, чтобы я его загадал - введи /wordle new  🤡"
			break
		}
		if err == wordle.WordNotExistsErr {
			content = "Такого слова не существует, меня не проведешь 🤡"
			break
		}
		if err != nil {
			content = "😭😭😭 Что-то пошло не так, ИЗВИНИТЕ 😭😭😭"
			break
		}

		if hs.wordleService.IsGuessed() {
			content = ">>> 🔤 WORDLE 🔤\n" + answer + fmt.Sprintf("\nСлово: %v \n🎉🎉🎉 %v, ТЫ УГАДАЛ СЛОВО 🎉🎉🎉\n🤡 Потребовалось попыток: %v 🤡",
				strings.ToUpper(word), i.Member.Mention(), hs.wordleService.GetTries())
			break
		}

		content = ">>> 🔤 WORDLE 🔤\n" + answer + fmt.Sprintf("\nСлово: %v\n🤡🤡🤡 %v, ТЫ НЕ УГАДАЛ СЛОВО 🤡🤡🤡",
			strings.ToUpper(word), i.Member.Mention())

	default:
		content = "🤡 Я не знаю такой команды 🥲"

	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})

	if err != nil {
		log.Printf("couldn't respond to 'wordle new' command, got error: %v", err.Error())
	}
}
