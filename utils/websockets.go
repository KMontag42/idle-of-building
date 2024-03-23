package utils

import (
  "golang.org/x/net/websocket"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func EmitMessage(ws *websocket.Conn, messageType string, messageData string) error {
  message := Message{Type: messageType, Data: messageData}
  // build the "websocket-message" template with the message data
  rendered := message.Data
  err := websocket.Message.Send(ws, rendered)
  if err != nil {
    return err
  }
  return nil
}
