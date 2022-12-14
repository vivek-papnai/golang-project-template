package models

import (
	"github.com/volatiletech/null"
	"time"
)

type User struct {
	ID         string    `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	Email      string    `db:"email" json:"email"`
	CreatedAt  time.Time `db:"created_at" json:"createdAt"`
	ArchivedAt null.Time `db:"archived_at" json:"archivedAt"`
}

const ActiveUser string = "active_user"

type WSMessageType string

const (
	WSMessageTypeChatRoom    WSMessageType = "chat_room"
	WSMessageTypeChatMessage WSMessageType = "chat"
)

type Message struct {
	Type WSMessageType `json:"type"`
	Data interface{}   `json:"data"`
}

type UserContext struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type ChatMessageInfo struct {
	Data    null.String `json:"data"`
	UserIDs []int       `json:"userIds"`
}

type PublishMessageData struct {
	Message       Message `json:"message"`
	SendToUserIDs []int   `json:"userIds"`
}
