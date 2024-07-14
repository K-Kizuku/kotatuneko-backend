package websocket

import (
	"context"
	"net/http"

	"github.com/K-Kizuku/kotatuneko-backend/internal/app/presentation/switcher"
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/service"
	"github.com/gorilla/websocket"
)

type IWSHandler interface {
	Start(ctx context.Context, w http.ResponseWriter, r *http.Request, switcher switcher.ISwitcher)
}

type WSHandler struct {
	msgSender service.IMessageSender
	physics   switcher.ISwitcher
}

func NewWSHandler(sender service.IMessageSender, physics switcher.ISwitcher) IWSHandler {
	return &WSHandler{
		msgSender: sender,
		physics:   physics,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (ws *WSHandler) Start(ctx context.Context, w http.ResponseWriter, r *http.Request, switcher switcher.ISwitcher) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	if err := switcher.Switch(ctx, conn); err != nil {
		return
	}
}
