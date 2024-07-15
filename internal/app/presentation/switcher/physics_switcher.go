package switcher

import (
	"context"
	"encoding/json"

	"github.com/K-Kizuku/kotatuneko-backend/internal/app/application/service"
	"github.com/K-Kizuku/kotatuneko-protobuf/gen/game/rpc"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type IPhysicsSwitcher interface {
	ISwitcher
}

type PhysicsSwitcher struct {
	physicsService service.IRoomObjectService
}

func NewPhysicsSwitcher(physicsService service.IRoomObjectService) IPhysicsSwitcher {
	return &PhysicsSwitcher{
		physicsService: physicsService,
	}
}

func (s *PhysicsSwitcher) Switch(ctx context.Context, conn *websocket.Conn) error {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		switch messageType {
		case websocket.TextMessage:
			var msg any
			if err := json.Unmarshal(p, &msg); err != nil {
				return err
			}
		case websocket.BinaryMessage:
			var msg rpc.PhysicsRequest
			if err := proto.Unmarshal(p, &msg); err != nil {
				return err
			}

			if err := s.physicsService.Calculate(ctx, "roomID"); err != nil {
				return err
			}
		default:
			return nil
		}

	}
}
