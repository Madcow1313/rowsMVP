package rowsUi

import (
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Data struct {
	Projects template.HTML
}

func InitAstilectron() {
	a, _ := astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName:            "rows",
		VersionAstilectron: "0.30.0",
		VersionElectron:    "25.2.0",
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

func InitMainWindow(a *astilectron.Astilectron) {
	files := FindAllProjectFiles()
	tFile, _ := os.Open("static/index.tmpl")
	defer tFile.Close()
	str, _ := io.ReadAll(tFile)
	t, _ := template.New("Projects").Parse(string(str))
	var tStrings string
	optStringFirst := `<option value=`
	optStringSecond := `</option>`
	appView, _ := os.OpenFile("index.html", os.O_CREATE|os.O_TRUNC, 0777)
	defer appView.Close()
	for _, val := range files {
		tStrings += optStringFirst + "\"" + val + "\">" + val + optStringSecond + "\n"
	}
	var d Data
	d.Projects = template.HTML(tStrings)
	err := t.Execute(appView, d)
	fmt.Println(err)
	w, _ := a.NewWindow("index.html", &astilectron.WindowOptions{
		Center:   astikit.BoolPtr(true),
		Height:   astikit.IntPtr(720),
		Width:    astikit.IntPtr(480),
		Show:     astikit.BoolPtr(true),
		Closable: astikit.BoolPtr(true),
	})
	w.Create()
	w.On(astilectron.EventNameWindowEventClosed, func(e astilectron.Event) (deleteListener bool) {
		w.Close()
		a.Stop()
		a.Close()
		return
	})
}
