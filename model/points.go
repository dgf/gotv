package model

type Point struct{ X, Y int }
type Points []Point

func (p Points) Len() int {
	return len(p)
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Points) Less(i, j int) bool {
	if p[i].Y < p[j].Y {
		return true
	}
	if p[i].X < p[j].X {
		return true
	}
	return false
}

// returns neighbours of a point
func neighbours(s int, p Point) []Point {
	a := []Point{}
	for _, n := range []Point{
		{X: p.X - 1, Y: p.Y}, // left
		{X: p.X + 1, Y: p.Y}, // right
		{X: p.X, Y: p.Y - 1}, // top
		{X: p.X, Y: p.Y + 1}, // bottom
	} {
		if n.X < 1 || n.X > s || n.Y < 1 || n.Y > s {
			continue // out
		}
		a = append(a, n)
	}
	return a
}

// diff two point slices
func DiffPoints(A, B []Point) []Point {
	d := []Point{}
	aVals := map[Point]struct{}{}
	bVals := map[Point]struct{}{}

	for _, a := range A {
		aVals[a] = struct{}{}
	}

	for _, b := range B {
		bVals[b] = struct{}{}
	}

	for _, b := range A {
		if _, ok := bVals[b]; !ok {
			d = append(d, b)
		}
	}

	for _, a := range B {
		if _, ok := aVals[a]; !ok {
			d = append(d, a)
		}
	}

	return d
}
