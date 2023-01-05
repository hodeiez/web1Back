package gui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Init() {
	a := app.New()
	myWindow := a.NewWindow("Form Layout")
	input := widget.NewEntry()
	label1 := widget.NewLabel("Label 1")
	value1 := widget.NewLabel("Value")
	label2 := widget.NewLabel("Label 2")
	label3 := widget.NewLabel("Label 3")
	but := widget.NewButton("button", func() {})
	// value2 := widget.NewLabel("Something")
	grid := container.New(layout.NewFormLayout(), label1, value1, label2, input, label3, but)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
	/*
		save album album panel
		save track track panel
		save infocard panel
	*/
}
