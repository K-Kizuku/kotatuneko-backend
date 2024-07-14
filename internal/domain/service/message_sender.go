package service

import "context"

type IMessageSender interface {
	Send(ctx context.Context, to string, data interface{}) error
	Broadcast(ctx context.Context, ids []string, data interface{}) error
}
