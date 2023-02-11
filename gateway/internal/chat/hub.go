package chat


import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[int64]*Client

	// Inbound messages from the clients.
	broadcast chan C2SMessage

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan C2SMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[int64]*Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.uid] = client
			hlog.Info("uid=", client.uid, " login success")
			h.sendOldMessage(*client)
		case client := <-h.unregister:
			if _, ok := h.clients[client.uid]; ok {
				delete(h.clients, client.uid)
				close(client.send)
			}
		case msg := <-h.broadcast:
			h.chooseClient(msg)
		}
	}
}

func (h Hub) sendOldMessage(c Client){
	msgs, err := GetUnsendMessageList(context.Background(), c.uid)
	if err!=nil{
		sentErr(c, err.Error())
	}

	for _, msg := range(msgs){
		h.chooseClient(C2SMessage{
			User_id: msg.FromUserId,
			To_user_id: msg.ToUserId,
			Msg_content: msg.MsgContent,
		})
	}

	err = SetUnsendMessage(context.Background(), c.uid)
	if err != nil{
		sentErr(c, err.Error())
	}
}

func (h Hub) chooseClient(msg C2SMessage) {
	if client, ok := h.clients[msg.To_user_id]; ok {
		select {
		case client.send <- msg:
		default:
			close(client.send)
			delete(h.clients, client.uid)
		}
	} else {
		// 将用户信息存入
		sendClient := h.clients[msg.User_id]

		msgdao := Message{
			FromUserId: msg.User_id,
			ToUserId: msg.To_user_id,
			MsgContent: msg.Msg_content,
		}
		err := CreateMessage(context.Background(), &msgdao)

		if err != nil{
			sentErr(*sendClient, err.Error())
		}else{
			sentErr(*sendClient, "user not login, messages save in sql")
		}
	}
}

func sentErr(c Client, errString string){
	c.send <- C2SMessage{
		User_id:     0,
		To_user_id:  0,
		Msg_content: errString,
	}
}