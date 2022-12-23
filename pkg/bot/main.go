package main

import (
	"diogonr.com/bot-in-go/pkg/bot/handlers/commands"
	"log"
	"os"
	"os/signal"
	"syscall"

	"diogonr.com/bot-in-go/pkg/bot/handlers"
	"github.com/bwmarrin/discordgo"
)

func main() {
	// init discord bot
	bot, err := discordgo.New("Bot ...")

	if err != nil {
		panic(err)
	}

	// start bot
	err = bot.Open()
	if err != nil {
		panic(err)
	}

	bot.Identify.Intents = discordgo.IntentMessageContent
	log.Println("Bot is running")

	// set bot prefix

	orchestrator := handlers.NewOrchestrator()

	orchestrator.Middleware().AddConditions(
		handlers.NoBot,
		handlers.NoSelf,
		handlers.NoWebhook,
		handlers.HasCommandPrefix,
		handlers.HasCommand,
	)

	emptyMiddleware := handlers.NewMiddleware()

	orchestrator.AddCommand("info", commands.Info, emptyMiddleware)
	orchestrator.AddCommand("ping", commands.Ping, emptyMiddleware)

	bot.AddHandler(orchestrator.Handler)

	defer bot.Close()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	<-stop
	log.Println("Bot is shutting down...")
}
