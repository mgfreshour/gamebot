package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
	"github.com/mgfreshour/gamebot/chess"
	"regexp"
)



const slackToken string = os.Getenv("SLACK_TOKEN")

func main() {
	api := slack.New(slackToken)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(false)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	var game *chess.Game
	game = chess.NewGame()
	moveRegex := regexp.MustCompile(`gamebot\s+([A-H])([0-8])-([A-H])([0-8])`)

	Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			//fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
			// Ignore hello


			case *slack.MessageEvent:
				if (ev.Text == "gamebot start game") {
					game = chess.NewGame()
					rtm.SendMessage(rtm.NewOutgoingMessage(game.DisplaySlack(), ev.Channel))
				}
				if (ev.Text == "gamebot show") {
					rtm.SendMessage(rtm.NewOutgoingMessage(game.DisplaySlack(), ev.Channel))
				}
				if (game != nil && moveRegex.MatchString(ev.Text)) {
					matches := moveRegex.FindAllStringSubmatch(ev.Text, -1)
					err := game.Move(matches[0][2], matches[0][1], matches[0][4], matches[0][3])
					if err != nil {
						rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("%v", err), ev.Channel))
					} else {
						rtm.SendMessage(rtm.NewOutgoingMessage("Made Move", ev.Channel))
						rtm.SendMessage(rtm.NewOutgoingMessage(game.DisplaySlack(), ev.Channel))
					}
				}

			case *slack.PresenceChangeEvent:
			//	fmt.Printf("Presence Change: %v\n", ev)

			case *slack.LatencyReport:
			//	fmt.Printf("Current latency: %v\n", ev.Value)

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Println("Invalid credentials")
				break Loop

			case *slack.DisconnectedEvent:
				log.Println("Bot disconnected")


			case *slack.ConnectedEvent:
				log.Printf("Bot connecting, connection_count=%d\n", ev.ConnectionCount)

			default:

			// Ignore other events..
			 fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}
}