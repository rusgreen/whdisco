package wh

import "strings"

func GetMarkdownHyperlink(content, url string) string {
	url = strings.ReplaceAll(url, "\\)", "\\\\\\)")
	return "[" + content + "](" + url + ")"
}
