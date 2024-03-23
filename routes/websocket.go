package routes

import (
  "github.com/labstack/echo/v4"
  "golang.org/x/net/websocket"
)

func WebSockets(c echo.Context) error {
  websocket.Handler(func(ws *websocket.Conn) {
    defer ws.Close()
    for {
      msg := "yo"
      if err := websocket.Message.Send(ws, msg); err != nil {
        c.Logger().Error(err)
        break
      }
      break
    }
  }).ServeHTTP(c.Response(), c.Request())
  return nil
}
