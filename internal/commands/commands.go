package commands

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "привет",
		Description: "Поздороваться с ботом 👋🏻",
	},
	{
		Name:        "wordle",
		Description: "Игра Wordle",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "new",
				Description: "Начать новую игру в Wordle",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        "guess",
				Description: "Попробовать отгадать слово",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "слово",
						Description: "слово из шести букв",
						Required:    true,
					},
				},
			},
			{
				Name:        "help",
				Description: "Как играть в Wordle❓",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
		},
	},
}
