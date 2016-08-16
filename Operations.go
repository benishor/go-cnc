package cnc

import "fmt"
import "github.com/benishor/go-cnc/gcode"
import "math"

type CncOperation interface {
	GetGCode() []string
}

type CuttingType int

const (
	CUTTING_TYPE_INSIDE CuttingType = iota
	CUTTING_TYPE_OUTSIDE CuttingType = iota
	CUTTING_TYPE_ON_PATH CuttingType = iota
)

type MachineSettings struct {
	PlungeFeedrate   float64
	MovementFeedrate float64
	SafeZ            float64
	CuttingDepth     float64
	CuttingType      CuttingType
}

type profileOperation struct {
	path     Path
	tool     Tool
	settings MachineSettings
}

func offsetPointInside(p Point, pathCenter Point, toolRadius float64) Point {
	newPoint := p

	if p.X < pathCenter.X {
		newPoint.X += toolRadius
	} else if p.X > pathCenter.X {
		newPoint.X -= toolRadius
	}

	if p.Y < pathCenter.Y {
		newPoint.Y += toolRadius
	} else if p.Y > pathCenter.Y {
		newPoint.Y -= toolRadius
	}

	return newPoint
}

func offsetPointOutside(p Point, pathCenter Point, toolRadius float64) Point {
	newPoint := p

	if p.X < pathCenter.X {
		newPoint.X -= toolRadius
	} else if p.X > pathCenter.X {
		newPoint.X += toolRadius
	}

	if p.Y < pathCenter.Y {
		newPoint.Y -= toolRadius
	} else if p.Y > pathCenter.Y {
		newPoint.Y += toolRadius
	}

	return newPoint
}

func (p *profileOperation) GetGCode() [] string {
	result := [] string{fmt.Sprintf("(Executing profile operation with tool %s)", p.tool.Name)}

	//cuttingDepthPerPass := float64(p.tool.Diameter / 1.1) // good default. or should it be /2?
	cuttingDepthPerPass := 2.2
	howManyPasses := int(math.Ceil(math.Abs(p.settings.CuttingDepth) / cuttingDepthPerPass))

	result = append(result, fmt.Sprintf("(cuttingDepthPerPass is %f and we need %d passes)", cuttingDepthPerPass, howManyPasses))

	toolRadius := p.tool.Diameter / 2.0
	pathCenter := p.path.GetCenter()

	for pass, passDepth := 0, 0.0; pass < howManyPasses; pass++ {
		passDepth -= cuttingDepthPerPass
		if passDepth < p.settings.CuttingDepth {
			passDepth = p.settings.CuttingDepth
		}

		result = append(result, fmt.Sprintf("(Pass %d, cutting to depth %f)", pass, passDepth))
		result = append(result, gcode.RaiseToolTo(p.settings.SafeZ, p.settings.PlungeFeedrate))

		newPoint := p.path.Points[0]
		switch p.settings.CuttingType {
		case CUTTING_TYPE_INSIDE:
			newPoint = offsetPointInside(p.path.Points[0], pathCenter, toolRadius)
		case CUTTING_TYPE_OUTSIDE:
			newPoint = offsetPointOutside(p.path.Points[0], pathCenter, toolRadius)
		}
		result = append(result, gcode.PositionToolTo(
			newPoint.X,
			newPoint.Y))


		result = append(result, gcode.LowerToolTo(passDepth, p.settings.PlungeFeedrate))

		for i := 0; i < len(p.path.Points); i++ {
			targetPointIndex := (i + 1) % len(p.path.Points)
			targetPoint := p.path.Points[targetPointIndex]

			switch p.settings.CuttingType {
			case CUTTING_TYPE_INSIDE:
				targetPoint = offsetPointInside(targetPoint, pathCenter, toolRadius)
			case CUTTING_TYPE_OUTSIDE:
				targetPoint = offsetPointOutside(targetPoint, pathCenter, toolRadius)
			}

			result = append(result, gcode.MoveTo(
				targetPoint.X,
				targetPoint.Y))
		}
		result = append(result, gcode.RaiseToolTo(p.settings.SafeZ, p.settings.PlungeFeedrate))
	}

	return result
}

func NewProfileOperation(path Path, t Tool, settings MachineSettings) *profileOperation {
	return &profileOperation{path: path, tool: t, settings: settings}
}
