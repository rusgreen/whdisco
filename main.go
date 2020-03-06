package main

import (
	"fmt"
	"github.com/rusgreen/whdisco/wh"
)

const discordChannelId = "your discord channel id"
const discordToken = "your discord token"
const discordWhUrl = "https://discordapp.com/api/webhooks/" + discordChannelId + "/" + discordToken

const build = 1

func main() {
	webhook := wh.NewDiscordWebhook(discordWhUrl)
	webhook.SetStatusGreen()
	webhook.SetTitle("Hello discord webhooks")
	description := "This **message** is sent via *discord hooks*.\n" +
		"Do you know that discord supports [markdown syntax](https://www.markdownguide.org/basic-syntax) ?"
	webhook.SetDescription(description)
	webhook.SetFooter(fmt.Sprintf("whdisco build #%d", build))
	err := webhook.Send()
	if err != nil {
		fmt.Println(err)
	}
}
