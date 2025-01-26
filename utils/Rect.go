package utils

type Rect struct {
	Position Vector2
	Size     Vector2
}

func (r Rect) GetCenter() Vector2 {
	return Vector2{r.Position.X - r.Size.X/2, r.Position.Y - r.Size.Y/2}
}
func (r *Rect) SetCentralizedPosition(center Vector2) {
	r.Position = Vector2{center.X + r.Size.X/2, center.Y + r.Size.Y/2}
}

func (r Rect) GetIntersectionWith(other Rect) Intersection {
	// Calculate the half sizes
	halfWidthA := r.Size.X / 2
	halfHeightA := r.Size.Y / 2
	halfWidthB := other.Size.X / 2
	halfHeightB := other.Size.Y / 2

	// Calculate centers
	centerA := r.GetCenter()
	centerB := other.GetCenter()

	// Calculate current and minimum-non-intersecting distances between centers
	distanceX := centerA.X - centerB.X
	distanceY := centerA.Y - centerB.Y
	minDistanceX := halfWidthA + halfWidthB
	minDistanceY := halfHeightA + halfHeightB

	// If we are not intersecting at all, return (0, 0)
	if distanceX > minDistanceX || distanceX < -minDistanceX || distanceY > minDistanceY || distanceY < -minDistanceY {
		return Intersection{0, Vector2{0, 0}}
	}

	// Calculate and return intersection depths
	depthX := minDistanceX - distanceX
	depthY := minDistanceY - distanceY
	if depthX < depthY {
		return Intersection{depthX, Vector2{depthX, 0}}
	}
	return Intersection{depthY, Vector2{0, depthY}}
}
