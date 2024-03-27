package utils

import (
	"bytes"
	"html/template"

	"github.com/kmontag42/idle-of-building/types"
	"golang.org/x/net/websocket"
)

func EmitMessage(ws *websocket.Conn, messageType string, messageData string) error {
	message := types.Message{Type: messageType, Data: messageData}
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

func EmitExperience(ws *websocket.Conn, character *types.Character) error {
	// build the "websocket-message" template with the message data
	var data bytes.Buffer

	template := template.Must(template.ParseFiles("views/pobpaste.html.tmpl"))
	err := template.ExecuteTemplate(&data, "experience-oob", character)
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

func EmitLife(ws *websocket.Conn, character *types.Character) error {
	// build the "websocket-message" template with the message data
	var data bytes.Buffer

	template := template.Must(template.ParseFiles("views/pobpaste.html.tmpl"))
	err := template.ExecuteTemplate(&data, "life-oob", character)
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


