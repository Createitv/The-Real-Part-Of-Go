package main

import (
	"fmt"
	"html"
)

func main() {
	const s = `"https:\u002F\u002Fphoto.zastatic.com\u002Fimages\u002Fphoto\u002F332356\u002F1329422446\u002F26653332996863252.jpg"`
	fmt.Println(html.UnescapeString(s))
}
