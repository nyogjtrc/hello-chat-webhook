package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"google.golang.org/api/chat/v1"
)

var url = ""

func main() {

	//msg := newTextMessage("hello, world!")
	msg := newCardMessage("work hard bot", "work work")

	_, err := sendMessage(msg)
	if err != nil {
		panic(err.Error())
	}
}

func sendMessage(m *chat.Message) (*chat.Message, error) {
	payload, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))
	return nil, nil
}

func newTextMessage(t string) *chat.Message {
	return &chat.Message{
		Text: t,
	}
}

func newCardMessage(title string, message string) *chat.Message {
	msg := chat.Message{}
	msg.CardsV2 = append(msg.CardsV2, &chat.CardWithId{
		Card: &chat.GoogleAppsCardV1Card{
			Header: &chat.GoogleAppsCardV1CardHeader{
				Title:    "Bot Message",
				Subtitle: "auto generate message",
				ImageUrl: "https://developers.google.com/chat/images/quickstart-app-avatar.png",
			},
			Sections: []*chat.GoogleAppsCardV1Section{
				{
					Header: title,
					Widgets: []*chat.GoogleAppsCardV1Widget{
						{
							TextParagraph: &chat.GoogleAppsCardV1TextParagraph{
								Text: message,
							},
						},
					},
				},
			},
		},
	})
	return &msg
}
