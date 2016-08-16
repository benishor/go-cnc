package cnc

type Point struct {
	X, Y float64
}

type Path struct {
	Points [] Point
}

func (c *Path) Add(points ...Point) () {
	for _, p := range points {
		c.Points = append(c.Points, p)
	}
}

func (c *Path) Translate(x, y float64) () {
	for i := range c.Points {
		c.Points[i].X += x
		c.Points[i].Y += y
	}
}

func (c *Path) Transpose() () {
	for i := range c.Points {
		oldX := c.Points[i].X
		c.Points[i].X = c.Points[i].Y
		c.Points[i].Y = oldX
	}
}

func (c *Path) GetCenter() Point {
	result := Point{}
	for _, p := range c.Points {
		result.X += p.X
		result.Y += p.Y
	}

	result.X /= float64(len(c.Points))
	result.Y /= float64(len(c.Points))

	return result
}
