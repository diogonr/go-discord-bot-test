package commands

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	log.Println("Command ping called by ", m.Author.Username)
	// check difference in time from command call to now
	// send message with difference
	diff := time.Since(m.Timestamp)

	s.ChannelMessageSend(m.ChannelID, "Pong! "+diff.String())
}
