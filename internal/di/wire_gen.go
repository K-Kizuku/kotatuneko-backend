// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	service2 "github.com/K-Kizuku/kotatuneko-backend/internal/app/application/service"
	"github.com/K-Kizuku/kotatuneko-backend/internal/app/infrastructure"
	"github.com/K-Kizuku/kotatuneko-backend/internal/app/presentation"
	"github.com/K-Kizuku/kotatuneko-backend/internal/app/presentation/handler"
	"github.com/K-Kizuku/kotatuneko-backend/internal/app/presentation/switcher"
	"github.com/K-Kizuku/kotatuneko-backend/internal/app/presentation/websocket"
	"github.com/K-Kizuku/kotatuneko-backend/internal/cat"
	"github.com/K-Kizuku/kotatuneko-backend/internal/cat/repository"
	"github.com/K-Kizuku/kotatuneko-backend/internal/cat/service"
	"github.com/K-Kizuku/kotatuneko-backend/pkg/cache"
)

// Injectors from wire.go:

func InitHandler() *presentation.Root {
	client := cache.NewCacheClient()
	iRoomObjectRepository := infrastructure.NewRoomObjectRepository(client)
	iMessageSender := infrastructure.NewMsgSender()
	iObjectRepository := repository.NewObjectRepository()
	iNikukyuRepository := repository.NewHandRepository()
	iHandService := service.NewHand(iObjectRepository, iNikukyuRepository)
	iObjectService := service.NewObjectService(iObjectRepository, iNikukyuRepository)
	iCatService := cat.New(iHandService, iObjectService)
	iCatRepository := infrastructure.NewCat(iCatService)
	iRoomObjectService := service2.NewRoomObjectService(iRoomObjectRepository, iMessageSender, iCatRepository)
	iwsHandler := websocket.NewWSHandler(iRoomObjectService, iMessageSender)
	iPhysicsSwitcher := switcher.NewPhysicsSwitcher(iRoomObjectService, iMessageSender)
	physicsHandler := handler.NewPhysicsHandler(iRoomObjectService, iwsHandler, iPhysicsSwitcher)
	root := presentation.New(physicsHandler)
	return root
}
