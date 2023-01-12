package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Init() {
	a := app.New()
	myWindow := a.NewWindow("TEST")
	myWindow.SetMaster()

	a.Settings().SetTheme(&MyTheme{})
	// panels := []*fyne.Container{TrackPanel(myWindow)}
	panels := []func(w fyne.Window) *fyne.Container{TrackPanel, AlbumPanel}
	myWindow.SetContent(MainPanel(a, myWindow, panels))
	// myWindow.SetContent(TrackPanel(myWindow))
	myWindow.ShowAndRun()
	/*
		save album album panel
		save track track panel
		save infocard panel
	*/
}
