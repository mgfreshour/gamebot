package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mgfreshour/gamebot/chess"
	"github.com/nlopes/slack"
	"regexp"
)

var slackToken = os.Getenv("SLACK_TOKEN")
var moveRegex = regexp.MustCompile(`gamebot\s+([A-H])([0-8])-([A-H])([0-8])`)
var game *chess.Game
var slackClient *slack.Client
var rtm *slack.RTM

func handlMessage (ev *slack.MessageEvent) {
	if ev.Text == "gamebot start game" {
		game = chess.NewGame()
		slackClient.PostMessage(ev.Channel, game.DisplaySlack(), slack.PostMessageParameters{})
	}
	if ev.Text == "gamebot show" {
		slackClient.PostMessage(ev.Channel, game.DisplaySlack(), slack.PostMessageParameters{})
	}
	if game != nil && moveRegex.MatchString(ev.Text) {
		matches := moveRegex.FindAllStringSubmatch(ev.Text, -1)
		err := game.Move(matches[0][2], matches[0][1], matches[0][4], matches[0][3])
		if err != nil {
			slackClient.PostMessage(ev.Channel, fmt.Sprintf("%v", err), slack.PostMessageParameters{})
		} else {
			slackClient.PostMessage(ev.Channel, "Made Move", slack.PostMessageParameters{})
			slackClient.PostMessage(ev.Channel, game.DisplaySlack(), slack.PostMessageParameters{})
		}
	}
}

func main() {
	if slackToken == "" {
		panic("Missing SLACK_TOKEN")
	}
	slackClient = slack.New(slackToken)
	slackClient.SetDebug(true)

	rtm = slackClient.NewRTM()
	go rtm.ManageConnection()

Loop:
	for msg := range rtm.IncomingEvents {
		// select {
		// case msg := <-:
			//fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
			// Ignore hello

			case *slack.MessageEvent:
				handlMessage(ev)

			case *slack.PresenceChangeEvent:
			//	fmt.Printf("Presence Change: %v\n", ev)

			case *slack.LatencyReport:
			//	fmt.Printf("Current latency: %v\n", ev.Value)

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Println("Invalid credentials!!")
				break Loop

			case *slack.DisconnectedEvent:
				log.Println("Bot disconnected")

			case *slack.ConnectedEvent:
				log.Printf("Bot connecting, connection_count=%d\n", ev.ConnectionCount)

			default:

				// Ignore other events..
				fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		// }
	}
}
