package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mgfreshour/gamebot/chess"
	"github.com/nlopes/slack"
	"regexp"
	"strings"
)

var slackToken = os.Getenv("SLACK_TOKEN")
var moveRegex = regexp.MustCompile(`gamebot\s+([a-h])([0-8])-([a-h])([0-8])`)
var slackClient *slack.Client
var rtm *slack.RTM

func sayInChan(ch string, msg string) {
	//chanId, Ts, err
	_, _, err := slackClient.PostMessage(ch, msg, slack.PostMessageParameters{Username: "gamebot"})
	if err != nil {
		fmt.Println(err)
	}
}

func handlMessage(ev *slack.MessageEvent) {
	text := strings.ToLower(ev.Text)
	if text == "gamebot start game" {
		game := chess.CreateGame(ev.Channel)
		sayInChan(ev.Channel, game.DisplaySlack())
	}
	if text == "gamebot show" {
		game := chess.LoadGame(ev.Channel)
		sayInChan(ev.Channel, game.DisplaySlack())
	}
	if moveRegex.MatchString(text) {
		game := chess.LoadGame(ev.Channel)
		if game == nil {
			sayInChan(ev.Channel, "No game exists in this channel")
			return
		}
		matches := moveRegex.FindAllStringSubmatch(text, -1)
		err := game.Move(matches[0][1], matches[0][2], matches[0][3], matches[0][4])
		if err != nil {
			sayInChan(ev.Channel, fmt.Sprintf("%v", err))
		} else {
			sayInChan(ev.Channel, "Made Move")
			sayInChan(ev.Channel, game.DisplaySlack())
			chess.SaveGame(game, ev.Channel)
		}
	}
}

func main() {
	if slackToken == "" {
		panic("Missing SLACK_TOKEN")
	}
	slackClient = slack.New(slackToken)
	slackClient.SetDebug(false)

	rtm = slackClient.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {

		case *slack.DisconnectedEvent:
			log.Println("Bot disconnected")

		case *slack.ConnectedEvent:
			log.Printf("Bot connecting, connection_count=%d\n", ev.ConnectionCount)
			// TODO - cache channel and users

		case *slack.MessageEvent:
			handlMessage(ev)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			panic("Invalid credentials!!")

			//case *slack.HelloEvent:
			// Ignore hello

			//case *slack.PresenceChangeEvent:
			//	fmt.Printf("Presence Change: %v\n", ev)

			//case *slack.LatencyReport:
			//	fmt.Printf("Current latency: %v\n", ev.Value)

			//default:
			// Ignore other events..
			//fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
