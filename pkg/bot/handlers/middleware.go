package handlers

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

type Middleware struct {
	conditions []func(s *discordgo.Session, m *discordgo.MessageCreate) bool
}

func (m *Middleware) AddCondition(self func(s *discordgo.Session, m *discordgo.MessageCreate) bool) *Middleware {
	log.Println("Adding condition")
	m.conditions = append(m.conditions, self)
	return m
}

func (m *Middleware) AddConditions(conditions ...func(s *discordgo.Session, m *discordgo.MessageCreate) bool) *Middleware {
	for _, condition := range conditions {
		m.AddCondition(condition)
	}
	return m
}

func NewMiddleware() *Middleware {
	return &Middleware{
		conditions: []func(s *discordgo.Session, m *discordgo.MessageCreate) bool{},
	}
}

func NoSelf(s *discordgo.Session, m *discordgo.MessageCreate) bool {

	return m.Author.ID != s.State.User.ID
}

func NoBot(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return !m.Author.Bot
}

func NoWebhook(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return m.WebhookID == ""
}

func HasCommandPrefix(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return m.Content[0] == '!'
}

func HasCommand(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return len(m.Content) > 1
}

func HasArguments(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return len(m.Content) > 2
}
