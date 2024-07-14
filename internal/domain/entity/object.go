package entity

type Object struct {
	ID       string
	Layer    int32
	Kinds    ObjectKind
	State    ObjectState
	Position Vector3
	Size     Vector3
}

type ObjectKind int32

const (
	ObjectKind_UNKNOWN ObjectKind = iota
	ObjectKind_RECTANGULAR
)

type ObjectState int32

const (
	ObjectState_UNKNOWN ObjectState = iota
	ObjectState_STAYING
	ObjectState_MOVING
	ObjectState_FORCING
)
