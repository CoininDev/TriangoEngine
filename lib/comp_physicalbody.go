package lib

type PhysicalBody struct {
	IsActive bool
	e        *Entity
	shape    Rect
	onFloor  bool
}

func NewPhysicalBody(size Vector2) *PhysicalBody {
	return &PhysicalBody{true, nil, Rect{Vector2{0, 0}, size}, false}
}

func (b *PhysicalBody) Start(e *Entity) { b.e = e }
func (b *PhysicalBody) GetType() string { return "PhysicalBody" }
func (b *PhysicalBody) Tick() {
	b.shape.Position = b.e.Position
	b.onFloor = false
	for _, col := range b.e.Game.Entities {
		// Ignorar colisão com a própria entidade
		if &col == b.e {
			continue
		}
		body := col.GetComponent("PhysicalBody").(*PhysicalBody)
		if body != nil {
			shape := body.shape
			b.onFloor = checkCollisionAtBottom(b.shape, shape)
		}
	}
}
func (b *PhysicalBody) End() {}

func checkCollisionAtBottom(shapeA, shapeB Rect) bool {
	// Verifica se shapeA está colidindo com shapeB na parte inferior
	return shapeA.Position.Y+shapeA.Size.Y == shapeB.Position.Y &&
		shapeA.Position.X < shapeB.Position.X+shapeB.Size.X &&
		shapeA.Position.X+shapeA.Size.X > shapeB.Position.X
}
