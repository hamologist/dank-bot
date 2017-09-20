package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/nlopes/slack"
)

var (
	oAuthToken string
)

func main() {
	api := slack.New(oAuthToken)
	validCharacterReg := regexp.MustCompile("[^a-z\\s]+")
	emojiFilterReg := regexp.MustCompile(":[a-z0-9+\\-_]+?:((:[a-z0-9+-_]+?:)+)?")
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.MessageEvent:
			text := validCharacterReg.ReplaceAllString(
				emojiFilterReg.ReplaceAllString(
					strings.ToLower(ev.Msg.Text), "",
				), "",
			)
			tokens := strings.Split(text, " ")

			for _, token := range tokens {
				if token == "lul" {
					rtm.SendMessage(rtm.NewOutgoingMessage("lol*", ev.Msg.Channel))
					break
				}
			}

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		default:
		}
	}
}

func init() {
	oAuthToken = os.Getenv("SLACK_OAUTH")

	if oAuthToken == "" {
		fmt.Print(
			"No slack OAuth token was provided.\n" +
				"Please consider providing an OAuth token using the \"SLACK_OAUTH\" env variable.\n",
		)
	}
}
