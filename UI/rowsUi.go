package rowsUi

import (
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"log"
	"os"
)

func InitAstilectron() {
	a, _ := astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName:            "rows",
		VersionAstilectron: "0.30.0",
		VersionElectron:    "25.2.0",
	})
	defer a.Close()
	a.Start()
	InitWindow(a)
	a.Wait()
}

func InitWindow(a *astilectron.Astilectron) {
	w, _ := a.NewWindow("http://127.0.0.1:8000", &astilectron.WindowOptions{
		Center:   astikit.BoolPtr(true),
		Height:   astikit.IntPtr(720),
		Width:    astikit.IntPtr(480),
		Show:     astikit.BoolPtr(true),
		Closable: astikit.BoolPtr(true),
	})
	w.Create()
	w.On(astilectron.EventNameWindowEventClosed, func(e astilectron.Event) (deleteListener bool) {
		//w.Close()
		//a.Quit()
		a.Stop()
		a.Close()
		return
	})
}
