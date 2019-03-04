package controllers

import (
	"github.com/gorilla/websocket"
	"github.com/astaxie/beego/logs"
	"LogView/models"
	"net/http"
)



type WebSocketController struct {
	BaseController
}

func (c *WebSocketController) Prepare() {
	c.BaseController.Prepare()
	c.checkLogin()
}

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan models.Message)
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  10240,
	WriteBufferSize: 10240,
	CheckOrigin: func(r *http.Request) bool { return true }}

func (c *WebSocketController) WsHandler() {
	//将http协议升级成websocket协议
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		logs.Error("http连接错误！", err)
		return
	}
	//defer ws.Close()
	clients[ws] = true
	for {
		var msg models.Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			logs.Error("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				logs.Error("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
func init() {
	go handleMessages()
}





