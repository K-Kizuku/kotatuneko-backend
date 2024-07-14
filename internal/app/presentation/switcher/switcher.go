package switcher

import (
	"context"

	"github.com/gorilla/websocket"
)

type ISwitcher interface {
	Switch(ctx context.Context, conn *websocket.Conn) error
}
