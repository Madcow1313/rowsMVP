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
	Number         string `json:"number,omitempty"`
	Date           string
	AbsoluteHeight float64
	HasWater       bool
	WaterDepths    []float64
	Depths         map[float64]Element `json:"depths,omitempty"`
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
