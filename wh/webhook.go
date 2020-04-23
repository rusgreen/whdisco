package wh

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty"
)

type DiscordWebhook struct {
	WebhookUrl string
	Obj        map[string]interface{}
	Embed      map[string]interface{}
	Fields     []map[string]interface{}
}

func NewDiscordWebhook(discordWhUrl string) *DiscordWebhook {
	result := &DiscordWebhook{}
	result.WebhookUrl = discordWhUrl
	result.Obj = make(map[string]interface{})
	result.Embed = make(map[string]interface{})
	result.Fields = make([]map[string]interface{}, 0)
	return result
}

func (d *DiscordWebhook) SetTitle(title string) *DiscordWebhook {
	d.Embed["title"] = title
	return d
}

func (d *DiscordWebhook) SetUrl(url string) *DiscordWebhook {
	d.Embed["url"] = url
	return d
}

func (d *DiscordWebhook) SetStatusGreen() *DiscordWebhook {
	d.Embed["color"] = COLOR_GREEN
	return d
}

func (d *DiscordWebhook) SetStatusRed() *DiscordWebhook {
	d.Embed["color"] = COLOR_RED
	return d
}

func (d *DiscordWebhook) SetStatusYellow() *DiscordWebhook {
	d.Embed["color"] = COLOR_YELLOW
	return d
}

func (d *DiscordWebhook) SetStatusGrey() *DiscordWebhook {
	d.Embed["color"] = COLOR_GREY
	return d
}

func (d *DiscordWebhook) SetDescription(description string) *DiscordWebhook {
	d.Embed["description"] = description
	return d
}

func (d *DiscordWebhook) SetContent(content string) *DiscordWebhook {
	d.Embed["description"] = content
	return d
}

func (d *DiscordWebhook) SetImage(url string) *DiscordWebhook {
	image := make(map[string]string)
	image["url"] = url
	d.Embed["image"] = image
	return d
}

func (d *DiscordWebhook) SetThumbnail(url string) *DiscordWebhook {
	thumbnail := make(map[string]string)
	thumbnail["url"] = url
	d.Embed["thumbnail"] = thumbnail
	return d
}

func (d *DiscordWebhook) AddField(name, value string) *DiscordWebhook {
	field := make(map[string]interface{})
	field["name"] = name
	field["value"] = value
	d.Fields = append(d.Fields, field)
	return d
}

func (d *DiscordWebhook) SetFooter(text string) *DiscordWebhook {
	footer := make(map[string]string)
	footer["text"] = text
	d.Embed["footer"] = footer
	return d
}

func (d *DiscordWebhook) Send() error {
	d.Embed["fields"] = d.Fields
	jsonBytes, err := json.Marshal(d.Embed)
	if err != nil {
		return err
	}
	if len(jsonBytes) > 6000 {
		return errors.New(fmt.Sprintf("Embed object larger than the limit (%d>6000).", len(jsonBytes)))
	}
	embeds := make([]map[string]interface{}, 1)
	embeds[0] = d.Embed
	d.Obj["embeds"] = embeds

	client := resty.New()
	response, err := client.R().
		SetHeader("Content-Type", "application/json;charset=utf-8").
		SetBody(d.Obj).
		Post(d.WebhookUrl)
	if err != nil {
		return err
	}
	if response.StatusCode() < 200 || response.StatusCode() >= 300 {
		return errors.New(string(response.Body()))
	}
	return nil
}
