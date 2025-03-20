package part2d

type Particle struct {
	Position Vector2D `json:"position"`
	Velocity Vector2D `json:"velocity"`
	Mass     float64  `json:"mass"`
}

type Vector2D struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
