package handlers

import (
	"clownbot/internal/services/wordle"
	"github.com/bwmarrin/discordgo"
)

type HandlerService struct {
	wordleService wordle.WordleService
}

func New(wordleService wordle.WordleService) *HandlerService {
	return &HandlerService{
		wordleService: wordleService,
	}
}

func (hs *HandlerService) Register(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"привет": hs.HelloHandler,
		"wordle": hs.WordleHandler,
	}
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}
