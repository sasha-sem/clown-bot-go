package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

// HelloHandler - handler for simple hello command
func (hs *HandlerService) HelloHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("👋🏻 Привет, 🤡Клоун🤡! %v", i.Member.Mention()),
		},
	})

	if err != nil {
		log.Printf("couldn't respond to 'привет' command, got error: %v", err.Error())
	}
}
