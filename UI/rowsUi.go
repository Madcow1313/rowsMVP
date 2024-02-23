package rowsUi

import (
	"encoding/json"
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"rowsMVP/DescriptionReader"
	"rowsMVP/Drawer"
	"strings"
)

type Data struct {
	Projects template.HTML
}

func InitAstilectron() {
	a, _ := astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName: "rows",
	})
	a.Start()
	InitMainWindow(a)
	a.Wait()
}

func FindAllProjectFiles() []string {
	dir, _ := os.ReadDir("./")
	files := make([]string, 0)
	for _, f := range dir {
		if filepath.Ext(f.Name()) == ".json" {
			fmt.Println(f.Name())
			files = append(files, f.Name())
		}
	}
	return files
}

func parseProjectTemplateOnCreation(d []DescriptionReader.Well) {
	type temp struct {
		Number template.HTML
	}
	templ, _ := os.Open("static/projectWindow.tmpl")
	defer templ.Close()
	s, _ := io.ReadAll(templ)
	parser, _ := template.New("Number").Parse(string(s))
	var data string
	optStringFirst := `<option value=`
	optStringSecond := `</option>`
	for _, well := range d {
		data += optStringFirst + "\"" + well.Number + "\">" + well.Number + optStringSecond + "\n"
	}
	var t temp
	t.Number = template.HTML(data)
	renderFile, _ := os.OpenFile("projectWindow.html", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0777)
	defer renderFile.Close()
	parser.Execute(renderFile, t)
}

func InitMainWindow(a *astilectron.Astilectron) {
	//files := FindAllProjectFiles()
	tFile, _ := os.Open("static/index.tmpl")
	defer tFile.Close()
	str, _ := io.ReadAll(tFile)
	t, _ := template.New("Projects").Parse(string(str))
	//var tStrings string
	//optStringFirst := `<option value=`
	//optStringSecond := `</option>`
	appView, _ := os.OpenFile("index.html", os.O_CREATE|os.O_TRUNC, 0777)
	//defer appView.Close()
	//for _, val := range files {
	//	tStrings += optStringFirst + "\"" + val + "\">" + val + optStringSecond + "\n"
	//}
	var d Data
	d.Projects = ""
	err := t.Execute(appView, d)
	fmt.Println(err)
	w, _ := a.NewWindow("index.html", &astilectron.WindowOptions{
		Center:         astikit.BoolPtr(true),
		Height:         astikit.IntPtr(720),
		Width:          astikit.IntPtr(960),
		Show:           astikit.BoolPtr(true),
		Closable:       astikit.BoolPtr(true),
		WebPreferences: &astilectron.WebPreferences{EnableRemoteModule: astikit.BoolPtr(true)},
	})
	w.Create()
	//w.OpenDevTools()
	w.OnMessage(func(m *astilectron.EventMessage) interface{} {
		var s string
		var found bool
		m.Unmarshal(&s)
		if s, found = strings.CutPrefix(s, "Project"); found {
			var we []DescriptionReader.Well
			file, _ := os.Open(s)
			defer file.Close()
			data, _ := io.ReadAll(file)
			err = json.Unmarshal(data, &we)
			parseProjectTemplateOnCreation(we)
			nw, _ := a.NewWindow("projectWindow.html", &astilectron.WindowOptions{
				Center:   astikit.BoolPtr(true),
				Height:   astikit.IntPtr(720),
				Width:    astikit.IntPtr(960),
				Show:     astikit.BoolPtr(true),
				Closable: astikit.BoolPtr(true),
			})
			nw.Create()
			nw.On(astilectron.EventNameWindowEventClosed, func(e astilectron.Event) (deleteListener bool) {
				w.Focus()
				nw.Close()
				nw.Destroy()
				return
			})
			nw.OnMessage(func(m *astilectron.EventMessage) interface{} {
				var s string
				m.Unmarshal(&s)
				switch s {
				case "add":

				}
				return nil
			})
			nw.Show()
			return nil
		} else if s, found = strings.CutPrefix(s, "File"); found {
			wells := DescriptionReader.ReadFile(s)
			var drawer Drawer.Drawer
			drawer.InitDrawer()
			drawer.SetWells(wells)
			drawer.DrawMain()
			filePath, _ := strings.CutSuffix(s, ".xlsx")
			drawer.Save(filePath + ".dxf")
		}
		return nil
	})

	w.On(astilectron.EventNameWindowEventClosed, func(e astilectron.Event) (deleteListener bool) {
		w.Close()
		a.Stop()
		a.Close()
		return
	})
}
