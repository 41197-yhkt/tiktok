package chat

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dns := "root:123456@tcp(127.0.0.1:6666)/tiktok_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dns),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})
	if err != nil {
		panic(err)
	}
	

}

type Message struct {
	Id         int
	FromUserId int64  `json:"from_user_id"`
	ToUserId   int64  `json:"to_user_id"`
	MsgContent string `json:"msg_content"`
	IsSend     bool   `json:"is_send"`
}

func (m *Message) TableName() string {
	return "messages"
}

func CreateMessage(ctx context.Context, msg *Message) error {
	if err := db.WithContext(ctx).Create(msg).Error; err != nil {
		return err
	}
	hlog.Info("create message in sql successful")
	return nil
}

// 找到和指定用户的消息记录(包括发送的消息和接受的消息)
func GetAllMessageList(ctx context.Context, from, to int64) ([]*Message, error) {
	var res []*Message
	if err := db.WithContext(ctx).Where("from_user_id = ? OR to_user_id = ?", from, to).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// 找到该用户未接受的消息
func GetUnsendMessageList(ctx context.Context, to int64) ([]*Message, error) {
	var res []*Message
	if err := db.WithContext(ctx).
		Where("to_user_id = ? AND is_send = 0", to).
		Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

// 设置该用户未接受的消息
func SetUnsendMessage(ctx context.Context, to int64) error {
	if err := db.WithContext(ctx).Model(&Message{}).
		Where("to_user_id =? AND is_send = false", to).
		Update("is_send", true).Error; err != nil {
		return err
	}
	return nil
}
