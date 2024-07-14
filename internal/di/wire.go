//go:build wireinject
// +build wireinject

package di

import (
	"github.com/K-Kizuku/kotatuneko-backend/internal/app/application/service"
	"github.com/K-Kizuku/kotatuneko-backend/internal/app/infrastructure"
	"github.com/K-Kizuku/kotatuneko-backend/internal/app/presentation"
	"github.com/K-Kizuku/kotatuneko-backend/internal/app/presentation/handler"
	"github.com/K-Kizuku/kotatuneko-backend/internal/app/presentation/switcher"
	"github.com/K-Kizuku/kotatuneko-backend/internal/app/presentation/websocket"
	"github.com/K-Kizuku/kotatuneko-backend/pkg/cache"
	"github.com/google/wire"
)

func InitHandler() *presentation.Root {
	wire.Build(
		cache.NewCacheClient,
		infrastructure.NewMsgSender,
		infrastructure.NewRoomObjectRepository,
		service.NewRoomObjectService,
		switcher.NewPhysicsSwitcher,
		websocket.NewWSHandler,
		handler.NewPhysicsHandler,
		presentation.New,
	)
	return &presentation.Root{}
}
