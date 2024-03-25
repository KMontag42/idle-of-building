package utils

import (
	"bytes"
	"html/template"

	"golang.org/x/net/websocket"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
	Map  string `json:"map"`
}

func EmitMessage(ws *websocket.Conn, messageType string, messageData string) error {
	message := Message{Type: messageType, Data: messageData}
	// build the "websocket-message" template with the message data
	var data bytes.Buffer

	template := template.Must(template.ParseFiles("views/battle.html.tmpl"))
	err := template.ExecuteTemplate(&data, "websocket-message", message)
	if err != nil {
		return err
	}

	rendered := data.String()

	err = websocket.Message.Send(ws, rendered)
	if err != nil {
		return err
	}
	return nil
}

