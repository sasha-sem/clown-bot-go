package main

import (
	"clownbot/internal/commands"
	"clownbot/internal/handlers"
	"clownbot/internal/services/wordle"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Getting bot token
	botToken, exists := os.LookupEnv("BOT_TOKEN")
	if !exists {
		log.Fatal("no authentication token found in .env")
		return
	}

	// Authenticating bot
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatalf("can't start authenticate bot, got error %v", err.Error())
		return
	}

	dictPath, exists := os.LookupEnv("WORDS_PATH")
	if !exists {
		log.Fatal("WORDS_PATH not found in .env")
		return
	}

	// Services
	wordleService, err := wordle.New(dictPath)
	if err != nil {
		log.Fatalf("Couldn't create words service, got error: %v", err)
		return
	}

	// Register handlers
	handlersServ := handlers.New(wordleService)
	dg.AddHandler(handlersServ.Register)
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})

	// Open connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		log.Fatalf("error opening connection, %v", err.Error())
		return
	}

	// Add commands
	for _, v := range commands.Commands {
		_, err = dg.ApplicationCommandCreate(dg.State.User.ID, "", v)
		if err != nil {
			log.Panicf("cann't create '%v' command: %v", v.Name, err)
		}
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Println("🤡🤡🤡 Clown Bot 🤡🤡🤡 is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	err = dg.Close()
	if err != nil {
		log.Fatalf("connection to bot was closed with error: %v", err.Error())
		return
	}
}
