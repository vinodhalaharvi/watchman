package main

import (
	"fmt"

	"github.com/bluele/slack"
)

const (
	token       = "xoxp-37472707681-37472707713-646594743329-7f6588688bbc053301ce1c7b27597b59"
	channelName = "admin"
)

func postToSlack(message string) {
	api := slack.New(token)
	err := api.ChatPostMessage(channelName, message, nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
