package handlers

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type HandlerFunc func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)

type HandlerWithMiddleware struct {
	handler    HandlerFunc
	middleware *Middleware
}

type Orchestrator struct {
	commandMap       map[string]HandlerWithMiddleware
	globalMiddleware *Middleware
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		commandMap:       make(map[string]HandlerWithMiddleware),
		globalMiddleware: NewMiddleware(),
	}
}

func (orchestrator *Orchestrator) Middleware() *Middleware {
	return orchestrator.globalMiddleware
}

func (orchestrator *Orchestrator) AddCommand(command string, handler HandlerFunc, middleware *Middleware) {
	log.Println("Adding command", command)
	orchestrator.commandMap[command] = HandlerWithMiddleware{
		handler:    handler,
		middleware: middleware,
	}
}

func (orchestrator *Orchestrator) Handler(session *discordgo.Session, message *discordgo.MessageCreate) {
	// run global middleware
	for _, condition := range orchestrator.globalMiddleware.conditions {
		if !condition(session, message) {
			return
		}
	}

	// parse command and arguments
	fragments := strings.Split(strings.ToLower(message.Content), " ")

	command := fragments[0][1:]
	arguments := fragments[1:]

	cmdFunc := orchestrator.commandMap[command]

	if cmdFunc.handler == nil {
		log.Println("Command not found")
		return
	}

	for _, condition := range cmdFunc.middleware.conditions {
		if !condition(session, message) {
			log.Println("Middleware condition failed")
			return
		}
	}

	cmdFunc.handler(session, message, arguments)
}
