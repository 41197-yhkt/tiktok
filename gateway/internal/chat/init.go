package chat

var ChatHub *Hub

func InitHub() {
	ChatHub= newHub()
	go ChatHub.run()
}

func Init(){
	InitDB()
	InitHub()
}
