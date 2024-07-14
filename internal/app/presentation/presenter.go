package presentation

import (
	"github.com/K-Kizuku/kotatuneko-backend/internal/app/presentation/handler"
)

type Root struct {
	PhysicsHandler *handler.PhysicsHandler
}

func New(physicsHandler *handler.PhysicsHandler) *Root {
	return &Root{
		PhysicsHandler: physicsHandler,
	}
}
