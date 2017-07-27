package model

// Point intersection on the board
type Point struct{ X, Y int }

// Points sortable point slice
type Points []Point

// Len is part of sort.Interface
func (p Points) Len() int {
	return len(p)
}

// Swap is part of sort.Interface
func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less is part of sort.Interface
func (p Points) Less(i, j int) bool {
	if p[i].X < p[j].X || p[i].Y < p[j].Y {
		return true
	}
	return false
}

// returns neighbours of a point
func neighbours(size int, point Point) (n []Point) {
	for _, p := range []Point{
		{X: point.X - 1, Y: point.Y}, // left
		{X: point.X + 1, Y: point.Y}, // right
		{X: point.X, Y: point.Y - 1}, // top
		{X: point.X, Y: point.Y + 1}, // bottom
	} {
		if p.X < 1 || p.X > size || p.Y < 1 || p.Y > size {
			continue // out
		}
		n = append(n, p)
	}
	return
}

// DiffPoints of two point slices
func DiffPoints(A, B []Point) (d []Point) {
	aPoints := map[Point]struct{}{}
	for _, a := range A {
		aPoints[a] = struct{}{}
	}

	bPoints := map[Point]struct{}{}
	for _, b := range B {
		bPoints[b] = struct{}{}
	}

	for _, a := range A { // a in B?
		if _, ok := bPoints[a]; !ok {
			d = append(d, a)
		}
	}

	for _, b := range B { // b in A?
		if _, ok := aPoints[b]; !ok {
			d = append(d, b)
		}
	}
	return
}
