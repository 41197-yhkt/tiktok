package chat

// Server to Client Message
type S2CMessage struct{
	From_user_id int64
    Msg_content string
}

// Client to Server message
type C2SMessage struct{
	User_id int64
    To_user_id int64
    Msg_content string
}