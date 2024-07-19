package service

import (
	"github.com/K-Kizuku/kotatuneko-backend/internal/cat/physics"
	"github.com/K-Kizuku/kotatuneko-backend/internal/cat/repository"
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
)

type IObjectService interface {
	GetObjectsSlice() []*entity.Nekojarashi
	GetObjectByObjID(key string) *entity.Nekojarashi
	CollideWithObj() map[string][]string
	ApplyForceToObj(obj1ID, obj2ID string)
}

type ObjectService struct {
	or repository.IObjectRepository
	nr repository.INikukyuRepository
}

func NewObjectService(or repository.IObjectRepository, nr repository.INikukyuRepository) IObjectService {
	return &ObjectService{
		or: or,
		nr: nr,
	}
}

func (os *ObjectService) GetObjectsSlice() []*entity.Nekojarashi {
	return os.or.GetObjectsSlice()
}

func (os *ObjectService) GetObjectByObjID(key string) *entity.Nekojarashi {
	return os.or.GetObjectByObjID(key)
}

// 全オブジェクトの衝突判定
func (os *ObjectService) CollideWithObj() map[string][]string {
	allObj := os.or.GetObjectsSlice()
	collidedObjIDs := make(map[string][]string, len(allObj))
	for i := 0; i < len(allObj); i++ {
		for j := 0; j <= i; j++ {
			if allObj[i].ID == allObj[j].ID {
				continue
			}
			if collided := physics.IsColliding(allObj[i].Position, allObj[j].Position); collided {
				collidedObjIDs[allObj[i].ID] = append(collidedObjIDs[allObj[i].ID], allObj[j].ID)
			}
		}
	}

	return collidedObjIDs
}

func (os *ObjectService) ApplyForceToObj(obj1ID, obj2ID string) {
	obj1 := os.or.GetObjectByObjID(obj1ID)
	obj2 := os.or.GetObjectByObjID(obj2ID)
	physics.CollidedVelocity(obj1, obj2)
	physics.UpdatePosition(obj1)
	physics.UpdatePosition(obj2)
}
