package Drawer

import (
	"github.com/yofu/dxf"
	"github.com/yofu/dxf/drawing"
	"github.com/yofu/dxf/entity"
	_ "github.com/yofu/dxf/table"
	dr "rowsMVP/DescriptionReader"
)

type Well dr.Well

type DrawerPosition struct {
	X, Y float64
}

type Drawer struct {
	fileName string
	drawer   *drawing.Drawing
	position DrawerPosition
	wells    []Well
}

func (d *Drawer) Wells() []Well {
	return d.wells
}

func (d *Drawer) SetWells(wells []Well) {
	d.wells = wells
}

func (d *Drawer) drawLine(startPos DrawerPosition, endPos DrawerPosition) {
	line := entity.NewLine()
	line.Start = []float64{startPos.X, startPos.Y, 0}
	line.End = []float64{endPos.X, endPos.Y, 0}
	d.drawer.AddEntity(line)
}

func (d *Drawer) putText(pos DrawerPosition, value string) {
	t := entity.NewText()
	t.Height = 2
	t.Coord1 = []float64{pos.X, pos.Y, 0}
	t.Value = value
	t.Rotation = 90
	t.WidthFactor = 1
	d.drawer.AddEntity(t)
}

// TODO:refactor this
func (d *Drawer) DrawTop() {
	endPos := d.position
	endPos.X += 175
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
	d.drawLine(DrawerPosition{X: d.position.X, Y: endPos.Y}, endPos)
	var i float64 = 0
	for x := d.position.X; x <= 40; x += 10 {
		d.position.X = x
		endPos.X = x
		d.putText(DrawerPosition{X: d.position.X + 5, Y: d.position.Y - 45}, top[i])
		i++
		d.drawLine(d.position, endPos)
	}
	r := []float64{25, 15, 15, 10, 70}
	for _, x := range r {
		d.position.X += x
		endPos.X = d.position.X
		d.putText(DrawerPosition{X: d.position.X + 5, Y: d.position.Y - 45}, top[i])
		i++
		d.drawLine(d.position, endPos)
	}

	endPos.Y = d.position.Y
	d.drawLine(d.position, endPos)
	d.drawer.Save()
}

func (d *Drawer) InitDrawer() {
	d.drawer = dxf.NewDrawing()
	d.drawer.Header().Version = "AC1024"

	d.drawer.SaveAs("test.dxf")
}
