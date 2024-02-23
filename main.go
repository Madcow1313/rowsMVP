package main

import (
	_ "github.com/asticode/go-astilectron"
	"github.com/yofu/dxf"
	"github.com/yofu/dxf/color"
	"github.com/yofu/dxf/entity"
	"github.com/yofu/dxf/table"
	"log"
	rowsUi "rowsMVP/UI"
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
	rowsUi.InitAstilectron()
}
