package gcode

import "fmt"

func RaiseToolTo(height float64, feedRate float64) string {
	return fmt.Sprintf("G00 Z%.6f F%.6f", height, feedRate)
}

func LowerToolTo(height float64, feedRate float64) string {
	return fmt.Sprintf("G01 Z%.6f F%.6f", height, feedRate)
}

func PositionToolTo(x float64, y float64) string {
	return fmt.Sprintf("G00 X%.6f Y%.6f", x, y)
}

func MoveTo(x float64, y float64) string {
	return fmt.Sprintf("G01 X%.6f Y%.6f", x, y)
}

func CircleAtCenterWithRadius(centerX, centerY, radius, depth, feedRate float64) ([] string) {
	// starting point
	x := centerX - radius
	y := centerY

	nextX := x + radius
	nextY := y + radius

	i := centerX - x
	j := centerY - y

	var result [] string
	result = append(result, fmt.Sprintf("G00 X%.6f Y%.6f", x, y))
	result = append(result, fmt.Sprintf("G01 Z%.6f F%.6f", depth, feedRate))
	result = append(result, fmt.Sprintf("G02 X%.6f Y%.6f I%.6f J%.6f F400", nextX, nextY, i, j))

	// advance
	x = nextX
	y = nextY
	nextX += radius
	nextY -= radius
	i = centerX - x
	j = centerY - y
	result = append(result, fmt.Sprintf("G02 X%.6f Y%.6f I%.6f J%.6f F400", nextX, nextY, i, j))

	// advance
	x = nextX
	y = nextY
	nextX -= radius
	nextY -= radius
	i = centerX - x
	j = centerY - y
	result = append(result, fmt.Sprintf("G02 X%.6f Y%.6f I%.6f J%.6f F400", nextX, nextY, i, j))

	// advance
	x = nextX
	y = nextY
	nextX -= radius
	nextY += radius
	i = centerX - x
	j = centerY - y
	result = append(result, fmt.Sprintf("G02 X%.6f Y%.6f I%.6f J%.6f F400", nextX, nextY, i, j))

	result = append(result, RaiseToolTo(4, feedRate))

	return result
}

func CircleWithRadius(x, y, radius float64) ([] string) {
	centerX := x + radius
	centerY := y

	nextX := x + radius
	nextY := y + radius

	i := centerX - x
	j := centerY - y

	var result [] string
	result = append(result, fmt.Sprintf("G02 X%.6f Y%.6f I%.6f J%.6f F400", nextX, nextY, i, j))

	// advance
	x = nextX
	y = nextY
	nextX += radius
	nextY -= radius
	i = centerX - x
	j = centerY - y
	result = append(result, fmt.Sprintf("G02 X%.6f Y%.6f I%.6f J%.6f F400", nextX, nextY, i, j))

	// advance
	x = nextX
	y = nextY
	nextX -= radius
	nextY -= radius
	i = centerX - x
	j = centerY - y
	result = append(result, fmt.Sprintf("G02 X%.6f Y%.6f I%.6f J%.6f F400", nextX, nextY, i, j))

	// advance
	x = nextX
	y = nextY
	nextX -= radius
	nextY += radius
	i = centerX - x
	j = centerY - y
	result = append(result, fmt.Sprintf("G02 X%.6f Y%.6f I%.6f J%.6f F400", nextX, nextY, i, j))

	return result
}
