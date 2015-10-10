package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func do(cmd string, rtm *slack.RTM, channel string) {
	// send docker output to slack
	go func() {
		out, err := exec.Command("sh", "-c", cmd).Output()

		if err != nil {
			rtm.SendMessage(rtm.NewOutgoingMessage(err.Error(), channel))
		} else {
			var lines []string
			lines = append(lines, "```")
			lines = append(lines, string(out[:]))
			lines = append(lines, "```")

			rtm.SendMessage(rtm.NewOutgoingMessage(strings.Join(lines, "\n"), channel))
		}
	}()
}

// interpret command sent to bot
func lu(cmd string, rtm *slack.RTM, channel string) error {
	do(cmd, rtm, channel)

	return nil
}

func main() {
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {

			case *slack.MessageEvent:
				messageText := ev.Text
				channel := ev.Channel
				fmt.Printf("Channel: %s\n", channel)

				// is it addressed to the bot?
				r, _ := regexp.Compile("^<@U0ASA381Z>:{0,1}")

				if r.MatchString(messageText) {
					cmd := r.ReplaceAllString(messageText, "")
					err := lu(cmd, rtm, channel)
					if err != nil {
						fmt.Println(err)
					}
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				// Ignore other events..
				// fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}
}
