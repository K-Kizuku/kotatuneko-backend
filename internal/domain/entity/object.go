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

func (o *Object) DeepCopy() *Object {
	return &Object{
		ID:    o.ID,
		Layer: o.Layer,
		Kinds: o.Kinds,
		State: o.State,
		Position: Vector3{
			X: o.Position.X,
			Y: o.Position.Y,
			Z: o.Position.Z,
		},
		Size: Vector3{
			X: o.Size.X,
			Y: o.Size.Y,
			Z: o.Size.Z,
		},
	}
}
