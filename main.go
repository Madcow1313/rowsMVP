package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	_ "github.com/asticode/go-astilectron"
	"github.com/yofu/dxf"
	"github.com/yofu/dxf/color"
	"github.com/yofu/dxf/entity"
	"github.com/yofu/dxf/table"
	"io"
	"log"
	"rowsMVP/DescriptionReader"
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
	//var drawer Drawer.Drawer
	//drawer.InitDrawer()

	var we []DescriptionReader.Well

	a := app.New()
	w := a.NewWindow("test")
	w.Resize(fyne.Size{
		Width:  720,
		Height: 480,
	})
	newBtn := widget.NewButton("Create", func() {
		fileDialog := dialog.NewFileOpen(func(closer fyne.URIReadCloser, e error) {
		}, w)
		fileDialog.Show()
	})
	w.SetContent(newBtn)
	openBtn := widget.NewButton("Open", func() {
		d := dialog.NewFileOpen(func(rc fyne.URIReadCloser, e error) {
			data, _ := io.ReadAll(rc)
			err := json.Unmarshal(data, &we)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(we)
		}, w)
		d.Show()
	})
	w.SetContent(openBtn)
	w.ShowAndRun()
	a.Quit()
	//drawer.SetWells([]DescriptionReader.Well{well})
	//drawer.DrawMain()
}
