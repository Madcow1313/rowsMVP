package DescriptionReader

import (
	"github.com/xuri/excelize/v2"
	"log"
)

type Element struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

type Project struct {
	Data []Well `json:"data"`
}

type Well struct {
	Number         string    `json:"number"`
	Date           string    `json:"date"`
	AbsoluteHeight float64   `json:"absoluteHeight"`
	HasWater       bool      `json:"hasWater"`
	WaterDepths    []float64 `json:"waterDepths"`
	RawDepths      []float64 `json:"depths"`
	Soil           []string  `json:"soil"`
	Identifiers    []string  `json:"identifiers"`
	Depths         map[float64]Element
}

func (w *Well) Elements() map[float64]Element {
	m := make(map[float64]Element)
	for i := range w.RawDepths {
		m[w.RawDepths[i]] = Element{
			Identifier: w.Identifiers[i],
			Name:       w.Soil[i],
		}
	}
	return m
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
