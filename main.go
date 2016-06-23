package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/huguesalary/slack-go"
)

func main() {
	client := slack.Client{}
	message := slack.Message{}
	flag.StringVar(&client.Url, "url", os.Getenv("SLACKP_URL"), "URL of the Slack webhook.")
	flag.StringVar(&message.Channel, "channel", os.Getenv("SLACKP_CHANNEL"), "Channel to post the message.")
	flag.StringVar(&message.Username, "user", os.Getenv("SLACKP_USER"), "Name of the user to post as.")
	flag.StringVar(&message.IconEmoji, "icon", os.Getenv("SLACKP_ICON"), "Icon of the message.")
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] TEXT\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		message.Text = strings.Join(args, " ")
	}

	err := client.SendMessage(&message)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
