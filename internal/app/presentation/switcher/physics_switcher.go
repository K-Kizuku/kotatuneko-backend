package switcher

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/K-Kizuku/kotatuneko-backend/internal/app/application/service"
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
	domainService "github.com/K-Kizuku/kotatuneko-backend/internal/domain/service"
	"github.com/K-Kizuku/kotatuneko-protobuf/gen/game/resources"
	"github.com/K-Kizuku/kotatuneko-protobuf/gen/game/rpc"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type IPhysicsSwitcher interface {
	ISwitcher
}

type PhysicsSwitcher struct {
	physicsService service.IRoomObjectService
	msgSender      domainService.IMessageSender
}

func NewPhysicsSwitcher(physicsService service.IRoomObjectService, msgSender domainService.IMessageSender) IPhysicsSwitcher {
	return &PhysicsSwitcher{
		physicsService: physicsService,
		msgSender:      msgSender,
	}
}

func (s *PhysicsSwitcher) Switch(ctx context.Context, conn *websocket.Conn) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Recovered from a panic: %v", err)
		}
	}()
	x := &rpc.PhysicsRequest{
		SenderId: "roomID",
		RoomId:   "roomID",
		Hands: &resources.Hand{
			UserId: "roomID",
			State:  1,
			CenterPosition: &resources.Vector3{
				X: 0,
				Y: 0,
				Z: 0,
			},
			ActionPosition: &resources.Vector3{
				X: 0,
				Y: 0,
				Z: 0,
			},
		},
	}
	b, err := proto.Marshal(x)
	if err != nil {
		fmt.Println("err: ", err)
		return err
	}
	if err := s.msgSender.Send(ctx, "roomID", b); err != nil {
		fmt.Println("err: ", err)
		return err
	}
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		fmt.Println("messageType: ", messageType)
		switch messageType {
		case websocket.TextMessage:
			var msg any
			if err := json.Unmarshal(p, &msg); err != nil {
				return err
			}
		case websocket.BinaryMessage:
			var msg rpc.PhysicsRequest
			if err := proto.Unmarshal(p, &msg); err != nil {
				fmt.Println("err: ", err)
				return err
			}
			// fmt.Println("msg: ", msg)
			hand := &entity.Hand{
				UserID: msg.SenderId,
				State:  entity.HandState(msg.Hands.State),
				CenterPosition: entity.Vector3{
					X: msg.Hands.CenterPosition.X,
					Y: msg.Hands.CenterPosition.Y,
					Z: msg.Hands.CenterPosition.Z,
				},
				ActionPosition: entity.Vector3{
					X: msg.Hands.ActionPosition.X,
					Y: msg.Hands.ActionPosition.Y,
					Z: msg.Hands.ActionPosition.Z,
				},
			}
			fmt.Println("000000000000")

			if err := s.physicsService.Calculate(ctx, msg.SenderId, msg.RoomId, hand); err != nil {
				return err
			}
		default:
			return nil
		}

	}
}
