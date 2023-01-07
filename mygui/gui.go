package gui

import (
	"fyne.io/fyne/v2/app"
)

func Init() {
	a := app.New()
	myWindow := a.NewWindow("TEST")
	myWindow.SetMaster()
	// panels := []*fyne.Container{TrackPanel(myWindow)}
	myWindow.SetContent(MainPanel(a, myWindow, TrackPanel))
	// myWindow.SetContent(TrackPanel(myWindow))
	myWindow.ShowAndRun()
	/*
		save album album panel
		save track track panel
		save infocard panel
	*/
}
