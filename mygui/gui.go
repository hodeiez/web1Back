package gui

import (
	"fyne.io/fyne/v2/app"
)

func Init() {
	a := app.New()
	myWindow := a.NewWindow("TEST")

	myWindow.SetContent(TrackPanel())
	myWindow.ShowAndRun()
	/*
		save album album panel
		save track track panel
		save infocard panel
	*/
}
