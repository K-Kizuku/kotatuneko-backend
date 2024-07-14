package entity

type Hand struct {
	UserID         string
	State          HandState
	CenterPosition Vector3
	ActionPosition Vector3
}

type HandState int32

const (
	HandState_UNKNOWN HandState = iota
	HandState_HOLDING
	HandState_OPENING
)
