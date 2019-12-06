package main

import (
	"fmt"

	"github.com/bluele/slack"
)

const (
	token       = "YOUR_SLACK_TOKEN"
	channelName = "YOUR_SLACK_CHANNEL"
)

func postToSlack(message string) {
	api := slack.New(token)
	err := api.ChatPostMessage(channelName, message, nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
