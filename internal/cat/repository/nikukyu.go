package repository

import (
	"sync"

	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
)

type INikukyuRepository interface {
	ModifyNikukyu(map[string]*entity.Nikukyu)
	ModifyNikukyuByUserID(userID string, hand *entity.Nikukyu)
	GetNikukyuByUserID(userID string) *entity.Nikukyu
	GetNikukyus() map[string]*entity.Nikukyu
	DeleteNikukyu(userID string)
	DeleteAllNikukyus()
	TransferHandToNikukyu(hand *entity.Hand) *entity.Nikukyu
}

type NikukyuRepository struct {
	mux   sync.RWMutex
	hands map[string]*entity.Nikukyu
}

func NewHandRepository() INikukyuRepository {
	return &NikukyuRepository{
		mux:   sync.RWMutex{},
		hands: make(map[string]*entity.Nikukyu),
	}
}

func (hr *NikukyuRepository) ModifyNikukyu(hands map[string]*entity.Nikukyu) {
	hr.mux.Lock()
	defer hr.mux.Unlock()
	hr.hands = hands
}

func (hr *NikukyuRepository) ModifyNikukyuByUserID(userID string, hand *entity.Nikukyu) {
	hr.mux.Lock()
	defer hr.mux.Unlock()
	hr.hands[userID] = hand
}

func (hr *NikukyuRepository) GetNikukyuByUserID(userID string) *entity.Nikukyu {
	hr.mux.RLock()
	defer hr.mux.RUnlock()
	if _, ok := hr.hands[userID]; !ok {
		nikukyu := &entity.Nikukyu{
			UserID: userID,
		}
		hr.addNikukyu(nikukyu)
		return nikukyu.DeepCopy()
	}
	return hr.hands[userID].DeepCopy()
}

func (hr *NikukyuRepository) GetNikukyus() map[string]*entity.Nikukyu {
	hr.mux.RLock()
	defer hr.mux.RUnlock()
	hands := make(map[string]*entity.Nikukyu, len(hr.hands))
	for k, v := range hr.hands {
		hands[k] = v.DeepCopy()
	}
	return hands
}

func (hr *NikukyuRepository) DeleteNikukyu(userID string) {
	hr.mux.Lock()
	defer hr.mux.Unlock()
	delete(hr.hands, userID)
}

func (hr *NikukyuRepository) DeleteAllNikukyus() {
	hr.mux.Lock()
	defer hr.mux.Unlock()
	hr.hands = make(map[string]*entity.Nikukyu)
}

func (hr *NikukyuRepository) TransferHandToNikukyu(hand *entity.Hand) *entity.Nikukyu {
	nikukyu := hr.GetNikukyuByUserID(hand.UserID)
	return &entity.Nikukyu{
		UserID:             hand.UserID,
		State:              hand.State,
		CenterPosition:     hand.CenterPosition,
		ActionPosition:     hand.ActionPosition,
		Velocity:           nikukyu.Velocity,
		PrevCenterPosition: nikukyu.ActionPosition,
		PrevActionPosition: nikukyu.PrevActionPosition,
		PrevVelocity:       nikukyu.PrevVelocity,
	}
}

func (hr *NikukyuRepository) addNikukyu(nikukyu *entity.Nikukyu) {
	hr.mux.Lock()
	defer hr.mux.Unlock()
	hr.hands[nikukyu.UserID] = nikukyu
}
