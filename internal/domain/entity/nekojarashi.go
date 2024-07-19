package entity

type Nekojarashi struct {
	ID           string
	Layer        int32
	Kinds        ObjectKind
	State        ObjectState
	Position     Vector3
	Size         Vector3
	Acceleration Vector3
	Velocity     Vector3
	Mass         float64
	PrevPosition Vector3     // 1フレーム前の座標
	PrevState    ObjectState // 1フレーム前の状態
	PrevVelocity Vector3     // 1フレーム前の速度
}

func (n *Nekojarashi) DeepCopy() *Nekojarashi {
	return &Nekojarashi{
		ID:    n.ID,
		Layer: n.Layer,
		Kinds: n.Kinds,
		State: n.State,
		Position: Vector3{
			X: n.Position.X,
			Y: n.Position.Y,
			Z: n.Position.Z,
		},
		Size: Vector3{
			X: n.Size.X,
			Y: n.Size.Y,
			Z: n.Size.Z,
		},
		Acceleration: Vector3{
			X: n.Acceleration.X,
			Y: n.Acceleration.Y,
			Z: n.Acceleration.Z,
		},
		Velocity: Vector3{
			X: n.Velocity.X,
			Y: n.Velocity.Y,
			Z: n.Velocity.Z,
		},
		Mass: n.Mass,
		PrevPosition: Vector3{
			X: n.PrevPosition.X,
			Y: n.PrevPosition.Y,
			Z: n.PrevPosition.Z,
		},
		PrevState: n.PrevState,
		PrevVelocity: Vector3{
			X: n.PrevVelocity.X,
			Y: n.PrevVelocity.Y,
			Z: n.PrevVelocity.Z,
		},
	}
}
