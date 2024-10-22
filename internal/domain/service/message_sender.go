package service

import (
	"context"

	"github.com/gorilla/websocket"
)

type IMessageSender interface {
	Send(ctx context.Context, to string, data interface{}) error
	Broadcast(ctx context.Context, roomID string, data interface{}) error
	Register(roomID, userID string, conn *websocket.Conn, err chan error)
}
