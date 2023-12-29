package DescriptionReader

import (
	"github.com/xuri/excelize/v2"
	"log"
)

type Element struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

type Well struct {
	Number         string              `json:"Number,omitempty"`
	Date           string              `json:"Date"`
	AbsoluteHeight float64             `json:"AbsoluteHeight,string"`
	HasWater       bool                `json:"HasWater,string"`
	WaterDepths    []float64           `json:"WaterDepths,omitempty"`
	Depths         map[float64]Element `json:"Depths,omitempty"`
}

func findBeginning(f *excelize.File) Well {
	var well Well

	//cols, _ := f.GetCols("Sheet1")

	return well
}

func ReadFile(path string) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Fatal("file opening error: ", err)
	}
	defer f.Close()
	findBeginning(f)

}
