package switcher

import (
	"context"
	"log"

	"github.com/K-Kizuku/kotatuneko-backend/internal/app/application/service"
	"github.com/gorilla/websocket"
)

type PhysicsSwitcher struct {
	physicsService service.IRoomObjectService
}

func NewPhysicsSwitcher(physicsService service.IRoomObjectService) ISwitcher {
	return &PhysicsSwitcher{
		physicsService: physicsService,
	}
}

func (s *PhysicsSwitcher) Switch(ctx context.Context, conn *websocket.Conn) error {
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return err
		}
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
