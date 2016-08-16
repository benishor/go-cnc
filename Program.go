package cnc

import "os"
import "bufio"
import "strings"

type Program struct {
	operations [] CncOperation
}

func (p *Program) Add(op CncOperation) {
	p.operations = append(p.operations, op)
}

func (p *Program) getHeader() [] string {
	return [] string{
		"T1 (select tool)",
		"G17 (XY plane selection)",
		"G21 (millimeters)",
		"M03 S2000 (spindle on, clockwise, speed rpm)"}
}

func (p *Program) getFooter() [] string {
	return [] string{
		"M05 (stop spindle)",
		"M30 (program end)"}
}

func (p *Program) GetGCode() [] string {
	result := p.getHeader()
	for _, op := range p.operations {
		for _, line := range op.GetGCode() {
			result = append(result, line)
		}
	}

	for _, line := range p.getFooter() {
		result = append(result, line)
	}

	return result
}

func (p *Program) WriteTo(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(strings.Join(p.GetGCode(), "\n"))
	writer.Flush()
}