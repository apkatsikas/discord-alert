// Package main demonstrates a bare simple bot without a state cache. It logs
// all messages it sees into stderr.
package main

import "github.com/apkatsikas/discord-alert/alert"

func main() {
	// TODO - remove, add README, try using it in app
	alert.SendAlert("Alert content!")
}
