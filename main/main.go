// C02LZHR5Y83
package main

import (
	"log"

	"github.com/slack-go/slack"
)

func main() {

	OAUTH_TOKEN := "OAUTH_TOKEN" // Paste your bot user token here
	CHANNEL_ID := "CHANNEL_ID"   // Paste your channel id here

	api := slack.New(OAUTH_TOKEN)
	attachment := slack.Attachment{
		Pretext: "Pretext",
		Text:    "Hello from slap4msth!",
	}

	channelId, timestamp, err := api.PostMessage(
		CHANNEL_ID,
		slack.MsgOptionText("message text", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Fatalf("%s\n", err)
	}

	log.Printf("Message successfully sent to Channel %s at %s\n", channelId, timestamp)
}
