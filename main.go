package main

import (
	_ "github.com/asticode/go-astilectron"
	"github.com/yofu/dxf"
	"github.com/yofu/dxf/color"
	"github.com/yofu/dxf/entity"
	"github.com/yofu/dxf/table"
	"log"
	"rowsMVP/DescriptionReader"
	"rowsMVP/Drawer"
)

func drawLineTest() {
	drawing := dxf.NewDrawing()
	drawing.Header().LtScale = 1000
	drawing.AddLayer("test", color.White, table.LT_CONTINUOUS, true)
	drawing.ChangeLayer("test")
	drawing.Line(0, 0, 0, 1, 1, 1)
	line := entity.NewLine()
	line.Start = []float64{0, 0, 0}
	line.End = []float64{100, 100, 0}
	entities := drawing.Entities()
	entities.Add(line)
	err := drawing.SaveAs("1.dxf")
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	var drawer Drawer.Drawer
	drawer.InitDrawer()
	var well DescriptionReader.Well
	well.AbsoluteHeight = 100
	well.HasWater = true
	well.WaterDepths = []float64{1.5}
	well.Depths = make(map[float64]DescriptionReader.Element)
	well.Depths[1] = DescriptionReader.Element{
		Identifier: "2а",
		Name:       "суглинок тугопластичный",
	}
	well.Depths[2] = DescriptionReader.Element{
		Identifier: "2б",
		Name:       "суглинок мягкопластичный",
	}
	drawer.SetWells([]DescriptionReader.Well{well})
	drawer.DrawMain()
}
