package alert

import (
	"context"
	"fmt"
	"os"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
)

func SendAlert(alertContent string) error {
	if alertContent == "" {
		return fmt.Errorf("alert cannot be empty")
	}
	var token = os.Getenv("BOT_TOKEN")
	if token == "" {
		return fmt.Errorf("no $BOT_TOKEN given")
	}

	var channelString = os.Getenv("CHANNEL_ID")
	if channelString == "" {
		return fmt.Errorf("no $CHANNEL_ID given")
	}

	s := session.New("Bot " + token)

	channelSnowflake, err := discord.ParseSnowflake(channelString)
	if err != nil {
		return fmt.Errorf("failed to parse snowflake from string %v", channelString)
	}

	s.AddIntents(gateway.IntentGuildMessages)

	if err := s.Open(context.Background()); err != nil {
		return fmt.Errorf("failed to connect %v", err)
	}
	defer s.Close()

	channelID := discord.ChannelID(channelSnowflake)
	if msg, err := s.SendMessage(channelID, alertContent); err != nil {
		return fmt.Errorf("failed to send message %v - error was %v", msg, err)
	}
	return nil
}
