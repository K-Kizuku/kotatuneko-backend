package repository

import (
	"sync"

	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
)

type IObjectRepository interface {
	ModifyObjects(key string, value *entity.Nekojarashi)
	GetObjectByObjID(key string) *entity.Nekojarashi
	GetObjectsSlice() []*entity.Nekojarashi
	GetObjectsMap() map[string]*entity.Nekojarashi
	DeleteObject(key string)
}

type ObjectRepository struct {
	mux  sync.RWMutex
	objs map[string]*entity.Nekojarashi
}

func NewObjectRepository() IObjectRepository {
	return &ObjectRepository{
		mux:  sync.RWMutex{},
		objs: make(map[string]*entity.Nekojarashi),
	}
}

func (or *ObjectRepository) ModifyObjects(key string, value *entity.Nekojarashi) {
	or.mux.Lock()
	defer or.mux.Unlock()
	or.objs[key] = value
}

func (or *ObjectRepository) GetObjectByObjID(key string) *entity.Nekojarashi {
	or.mux.RLock()
	defer or.mux.RUnlock()
	return or.objs[key].DeepCopy()
}

func (or *ObjectRepository) GetObjectsSlice() []*entity.Nekojarashi {
	or.mux.RLock()
	defer or.mux.RUnlock()
	objs := make([]*entity.Nekojarashi, 0, len(or.objs))
	for _, obj := range or.objs {
		objs = append(objs, obj.DeepCopy())
	}
	return objs
}

func (or *ObjectRepository) GetObjectsMap() map[string]*entity.Nekojarashi {
	or.mux.RLock()
	defer or.mux.RUnlock()
	objMap := make(map[string]*entity.Nekojarashi, len(or.objs))
	for k, v := range or.objs {
		objMap[k] = v.DeepCopy()
	}
	return objMap
}

func (or *ObjectRepository) DeleteObject(key string) {
	or.mux.Lock()
	defer or.mux.Unlock()
	delete(or.objs, key)
}
