package constants

const (
	// 物理定数
	Gravity         = float64(9.8)  // 重力
	Restitution     = float64(0.5)  // 反発係数
	Friction        = float64(0.7)  // 摩擦係数
	Speed           = float64(10.0) // 最大速度
	Unit            = float64(1.0)  // 単位
	Minimum         = float64(0.0)  // 最小値
	KineticFriction = float64(0.5)  // 運動摩擦係数

	// ゲーム定数
	ObjectNum  = 5   // オブジェクト数
	PlayerNum  = 4   // プレイヤー数
	FrameRate  = 10  // フレームレート
	BlockSizeX = 1.0 // ブロックのX方向のサイズ
	BlockSizeY = 1.0 // ブロックのY方向のサイズ
	BlockSizeZ = 1.0 // ブロックのZ方向のサイズ
	BlockMass  = 1.0 // ブロックの質量
	TimeStep   = 0.1 // 時間刻み
)
