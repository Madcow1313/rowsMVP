package Drawer

import (
	"github.com/yofu/dxf"
	"github.com/yofu/dxf/drawing"
	"github.com/yofu/dxf/entity"
	_ "github.com/yofu/dxf/table"
	"rowsMVP/DescriptionReader"
	"strconv"
)

type Position struct {
	X, Y float64
}

type Drawer struct {
	fileName    string
	drawer      *drawing.Drawing
	position    Position
	description []DescriptionReader.Well
}

func (d *Drawer) Wells() []DescriptionReader.Well {
	return d.description
}

func (d *Drawer) SetWells(wells []DescriptionReader.Well) {
	d.description = wells
}

func (d *Drawer) drawLine(startPos Position, endPos Position) {
	line := entity.NewLine()
	line.Start = []float64{startPos.X, startPos.Y, 0}
	line.End = []float64{endPos.X, endPos.Y, 0}
	d.drawer.AddEntity(line)
}

func (d *Drawer) putText(pos Position, value string, rotation float64) {
	t := entity.NewText()
	t.Height = 2
	t.Coord1 = []float64{pos.X, pos.Y, 0}
	t.Value = value
	t.Rotation = rotation
	t.WidthFactor = 1
	d.drawer.AddEntity(t)
}

// TODO: add water and samples
func (d *Drawer) DrawMain() {
	wells := d.Wells()
	pos := d.position
	pos.Y = -50
	for _, w := range wells {
		prevDepth := float64(0)
		for depth, _ := range w.Depths {
			d.drawLine(Position{pos.X, pos.Y - prevDepth*10}, Position{pos.X, pos.Y - depth*10})
			d.drawLine(Position{pos.X, pos.Y - prevDepth*10}, Position{pos.X, pos.Y - depth*10})
			d.drawLine(Position{pos.X, pos.Y - depth*10}, Position{pos.X + 65, pos.Y - depth*10})
			d.drawLine(Position{pos.X + 95, pos.Y - depth*10}, Position{pos.X + 175, pos.Y - depth*10})

			d.putText(Position{pos.X + 5, pos.Y - prevDepth*10}, "", 0)
			d.putText(Position{pos.X + 10, pos.Y - depth*10 + 1}, strconv.FormatFloat(depth-prevDepth, 'f', 2, 64), 0)
			d.putText(Position{pos.X + 20, pos.Y - depth*10 + 1}, strconv.FormatFloat(depth, 'f', 2, 64), 0)
			d.putText(Position{pos.X + 30, pos.Y - depth*10 + 1},
				strconv.FormatFloat(w.AbsoluteHeight-depth-prevDepth, 'f', 2, 64), 0)

			d.putText(Position{pos.X + 85, pos.Y - prevDepth*10 - 5}, "", 0)
			d.putText(Position{pos.X + 100, pos.Y - prevDepth*10 - 5}, w.Depths[depth].Identifier, 0)
			d.putText(Position{pos.X + 110, pos.Y - prevDepth*10 - 5}, w.Depths[depth].Name, 0)

			prevDepth = depth
		}
		if w.HasWater {
			for _, h := range w.WaterDepths {
				d.putText(Position{pos.X + 65, pos.Y - h*10 - 5}, "", 0)
				d.drawLine(Position{pos.X + 65, pos.Y - h*10}, Position{
					pos.X + 80,
					pos.Y - h*10,
				})
			}
		}
		d.DrawTop(prevDepth)
	}
	d.drawer.Save()
}

// TODO:refactor this
func (d *Drawer) DrawTop(pd float64) {
	endPos := d.position
	endPos.X += 175
	startTemp := d.position
	top := map[float64]string{
		0: "Геоиндекс",
		1: "Мощность, м",
		2: "Глубина, м",
		3: "Абс. отметка, м",
		4: "Геолого-литологический разрез",
		5: "Сведения о воде",
		6: "Сведения о пробах",
		7: "ИГЭ",
		8: "Наименование и характеристика пород",
	}
	d.drawLine(d.position, endPos)
	endPos.Y = -50
	d.drawLine(Position{X: d.position.X, Y: endPos.Y}, endPos)
	endPos.Y = -50 - pd*10
	d.drawLine(Position{X: d.position.X, Y: endPos.Y}, endPos)
	var i float64 = 0
	for x := d.position.X; x <= 40; x += 10 {
		d.position.X = x
		endPos.X = x
		d.putText(Position{X: d.position.X + 5, Y: d.position.Y - 45}, top[i], 90)
		i++
		d.drawLine(d.position, endPos)
	}
	r := []float64{25, 15, 15, 10, 70}
	for _, x := range r {
		d.position.X += x
		endPos.X = d.position.X
		d.putText(Position{X: d.position.X + 5, Y: d.position.Y - 45}, top[i], 90)
		i++
		d.drawLine(d.position, endPos)
	}

	endPos.Y = d.position.Y
	d.drawLine(d.position, endPos)
	d.position = startTemp
	d.drawer.Save()
}

func (d *Drawer) InitDrawer() {
	d.drawer = dxf.NewDrawing()
	d.drawer.Header().Version = "AC1024"

	d.drawer.SaveAs("test.dxf")
}
