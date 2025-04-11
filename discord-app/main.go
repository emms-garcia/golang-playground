package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joomcode/errorx"
	"go.uber.org/zap"
)

func main() {
	logger := InitLogger(GetEnvironment())
	// create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + GetDiscordBotToken())
	if err != nil {
		zap.S().Panic(errorx.Decorate(err, "error creating Discord session"))
	}

	giphy := NewGiphyHandler(GetGiphyAPIKey())
	// register the messageCreate func as a callback for MessageCreate events
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// ignore all messages created by the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		message := strings.TrimSpace(m.Content)
		tokens := strings.Split(message, " ")
		if len(tokens) < 2 {
			return
		}
		if tokens[0] != fmt.Sprintf("<@%s>", s.State.User.ID) {
			return
		}

		command := tokens[1]
		commandParams := strings.Join(tokens[2:], " ")
		zap.S().Debugf("command: %s, Params: %s", command, commandParams)
		switch command {
		case "ping":
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		case "gif":
			gif, err := giphy.SearchFirst(commandParams)
			if err != nil {
				zap.S().Error(errorx.Decorate(err, "failed to fetch gif from Giphy"))
				return
			}

			s.ChannelMessageSend(m.ChannelID, gif.Url)
		}
	})

	// open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		zap.S().Panic(errorx.Decorate(err, "error opening connection"))
	}

	// wait here until CTRL-C or other term signal is received
	zap.S().Info("bot is now running. Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// cleanly close the Discord session before exiting
	dg.Close()
	// flush the logs before exiting
	logger.Sync()
}
