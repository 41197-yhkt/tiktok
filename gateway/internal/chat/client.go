package chat

import (
	"log"
	"strconv"
	"time"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan C2SMessage

	uid int64
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.

// Client在这个线程中读取socket中的内容并广播到Hub
// 大坑!居然在go里调指针
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		// 指针!!!
		var c2sMessage C2SMessage
		err := c.conn.ReadJSON(&c2sMessage)
		if err != nil {
			hlog.Info("read json err")
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		c.hub.broadcast <- c2sMessage
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		// server的client实例向真正的客户端发送信息
		case c2sMessage, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			s2cMessage := S2CMessage{
				From_user_id: c2sMessage.User_id,
				Msg_content:  c2sMessage.Msg_content,
			}
			//hlog.Info("s2cmessage=", s2cMessage)
			err := c.conn.WriteJSON(s2cMessage)
			if err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(ctx *app.RequestContext, hub *Hub) {

	uid, err := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	if err != nil {
		hlog.Info("uid not exit")
		return
	}
	hlog.Info("uid=", uid)

	err = upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
		// ATTENTION: 注意send缓冲区大小
		client := &Client{hub: hub,
			conn: conn,
			send: make(chan C2SMessage, 16),
			uid:  uid,
		}
		client.hub.register <- client

		go client.writePump()
		client.readPump()
	})
	if err != nil {
		log.Println(err)
	}
}

var upgrader = websocket.HertzUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
