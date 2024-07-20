package websocket

import (
	"context"
	"net/http"

	appService "github.com/K-Kizuku/kotatuneko-backend/internal/app/application/service"
	"github.com/K-Kizuku/kotatuneko-backend/internal/app/presentation/switcher"
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/service"
	"github.com/gorilla/websocket"
)

type IWSHandler interface {
	Start(ctx context.Context, w http.ResponseWriter, r *http.Request, switcher switcher.ISwitcher) error
}

type WSHandler struct {
	roomObjectService appService.IRoomObjectService
	msgSender         service.IMessageSender
}

func NewWSHandler(roomObjectService appService.IRoomObjectService, sender service.IMessageSender) IWSHandler {
	return &WSHandler{
		roomObjectService: roomObjectService,
		msgSender:         sender,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (ws *WSHandler) Start(ctx context.Context, w http.ResponseWriter, r *http.Request, switcher switcher.ISwitcher) error {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	errCh := make(chan error)
	ws.msgSender.Register("roomID", "roomID", conn, errCh)
	// defer conn.Close()

	go func() {
		if err := switcher.Switch(ctx, conn); err != nil {
			return
		}
	}()
	// defer close(errCh)
	return nil
}
