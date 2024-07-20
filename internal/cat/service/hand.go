package service

import (
	"github.com/K-Kizuku/kotatuneko-backend/internal/cat/constants"
	"github.com/K-Kizuku/kotatuneko-backend/internal/cat/physics"
	"github.com/K-Kizuku/kotatuneko-backend/internal/cat/repository"
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
)

type IHandService interface {
	CalculateHandForce(hand *entity.Nikukyu) *entity.Vector3 // 手の力の更新
	CollideWithObj(hand *entity.Nikukyu) *string             // ブロックとの衝突判定(衝突したブロックのIDを返す)
	TransferHandToNikukyu(hand *entity.Hand) *entity.Nikukyu
	ApplyForceToObj(id string, force *entity.Vector3) // ブロックに力を加える
}

type HandService struct {
	or repository.IObjectRepository
	nr repository.INikukyuRepository
}

func NewHand(or repository.IObjectRepository, nr repository.INikukyuRepository) IHandService {
	return &HandService{
		or: or,
		nr: nr,
	}
}

func (h *HandService) CalculateHandForce(hand *entity.Nikukyu) *entity.Vector3 {
	currHandVel := physics.CalculateVelocity(hand.PrevActionPosition, hand.ActionPosition, constants.TimeStep)
	handAcc := physics.CalculateAcceleration(hand.PrevVelocity, *currHandVel, constants.TimeStep)
	handForce := physics.CalculateForce(constants.BlockMass, *handAcc)
	return handForce
}

// 全Objectとの当たり判定を行い、当たったObjectのIDを返す
func (h *HandService) CollideWithObj(hand *entity.Nikukyu) *string {
	for _, v := range h.or.GetObjectsSlice() {
		if collided := physics.IsColliding(hand.ActionPosition, v.Position); collided {
			return &v.ID
		}
	}
	return nil
}

func (h *HandService) TransferHandToNikukyu(hand *entity.Hand) *entity.Nikukyu {
	nikukyu := h.nr.TransferHandToNikukyu(hand)
	return nikukyu
}

func (h *HandService) ApplyForceToObj(id string, force *entity.Vector3) {
	obj := h.or.GetObjectByObjID(id)
	physics.ApplyForce(obj, force)
	physics.ApplyFriction(obj)
	physics.UpdatePosition(obj)
}
