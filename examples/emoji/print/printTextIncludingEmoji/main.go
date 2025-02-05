package main

import (
	"fmt"
	"os"

	"go.m3o.com/emoji"
)

// Print text and renders the emojis with aliases e.g
// let's grab a :beer: becomes let's grab a 🍺
func main() {
	emojiService := emoji.NewEmojiService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := emojiService.Print(&emoji.PrintRequest{
		Text: "let's grab a :beer:",
	})
	fmt.Println(rsp, err)
}
