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
