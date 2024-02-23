package DescriptionReader

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"math"
	"strconv"
)

type Element struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

type Project struct {
	Data     []Well `json:"data"`
	FileName string
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

	rows, _ := f.GetRows("Sheet1")
	for _, r := range rows {
		for _, str := range r {
			fmt.Println(str)
		}
	}
	return well
}

//Hint: 0 - number, 1 - date, 2 - absolute height, 3 - rawDepths, 4 - id, 5 - not needed, 6 - not needed
// 7 - water, now not needed, 8,9,10 - not needed, 11- description, I have no idea what to do with that, 12 - add to 11?

func Walk(f *excelize.File) []Well {
	wells := make([]Well, 0)
	rows, _ := f.GetRows("Sheet1")
	for _, r := range rows {
		if len(r[0]) > 0 {
			absHeight, _ := strconv.ParseFloat(r[2], 32)
			rawDepths := make([]float64, 0)
			d, _ := strconv.ParseFloat(r[3], 32)
			rawDepths = append(rawDepths, math.Round(d*100)/100)
			ids := make([]string, 0)
			ids = append(ids, r[4])
			soil := make([]string, 0)
			soil = append(soil, r[len(r)-2])

			wells = append(wells, Well{
				Number:         r[0],
				Date:           r[1],
				AbsoluteHeight: math.Round(absHeight*100) / 100,
				RawDepths:      rawDepths,
				Identifiers:    ids,
				Soil:           soil,
			})
		} else {
			d, _ := strconv.ParseFloat(r[3], 32)
			if d > 0 {
				wells[len(wells)-1].RawDepths = append(wells[len(wells)-1].RawDepths, math.Round(d*100)/100)
			}
			if len(r[4]) > 0 {
				wells[len(wells)-1].Identifiers = append(wells[len(wells)-1].Identifiers, r[4])
			}
			wells[len(wells)-1].Soil = append(wells[len(wells)-1].Soil, r[len(r)-2])
		}
	}
	return wells
}

func ReadFile(path string) []Well {
	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Fatal("file opening error: ", err, "filepath: ", path)
	}
	defer f.Close()
	w := Walk(f)
	return w
}
