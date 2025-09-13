# discord-alert

discord-alert is a simple alerting library for Discord bots, written in Go.

## Creating a bot

0. Create a new Discord application
0. Take note of the bot token
0. Go to the OAuth2 tab
0. Add a redirect URL like https://discordapp.com/oauth2/authorize?&client_id=[CLIENTID]&scope=bot - replacing [CLIENTID] with your application's client ID
0. Visit the redirect URL in your browser to authorize the bot to a server you administrate

## Using the library

0. Set `BOT_TOKEN`
0. Set `CHANNEL_ID`

Example:

```
package main

import (
	"log"

	"github.com/apkatsikas/discord-alert/alert"
)

func main() {
	const content = "Alert content!"
	err := alert.SendAlert(content)
	if err != nil {
		log.Fatalf("failed to send %v - error was %v", content, err)
	}
}
```
