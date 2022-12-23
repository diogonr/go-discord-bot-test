package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Info(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	log.Printf("Command info called by %s", m.Author.Username)
	s.ChannelMessageSend(m.ChannelID, "This is a bot written in Go")
}
