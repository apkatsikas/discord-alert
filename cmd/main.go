// Package main demonstrates a bare simple bot without a state cache. It logs
// all messages it sees into stderr.
package main

import (
	"context"
	"log"
	"os"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
)

func main() {
	var token = os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatalln("No $BOT_TOKEN given.")
	}

	var channelString = os.Getenv("CHANNEL_ID")
	if channelString == "" {
		log.Fatalln("No $CHANNEL_ID given")
	}

	s := session.New("Bot " + token)

	channelID, err := discord.ParseSnowflake(channelString)

	if err != nil {
		log.Fatalf("failed to parse snowflake from string %v", channelString)
	}

	s.AddIntents(gateway.IntentGuildMessages)

	if err := s.Open(context.Background()); err != nil {
		log.Fatalln("Failed to connect:", err)
	}
	defer s.Close()

	msg, err := s.SendMessage(discord.ChannelID(channelID), "ALERT!")

	if err != nil {
		log.Fatalf("failed to send message %v - error was %v", msg, err)
	}
}
