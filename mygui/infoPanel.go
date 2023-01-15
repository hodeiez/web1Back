package gui

import (
	"hodei/web1/db"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func InfoPanel(w fyne.Window) *fyne.Container {
	selImageForInfo := widget.NewLabel("")
	selInfoType := widget.NewLabel("")

	titleLab := widget.NewLabel("Title")
	titleEntry := widget.NewEntry()
	descrLab := widget.NewLabel("Description")
	descrEntry := widget.NewEntry()
	yearLab := widget.NewLabel("Year")
	yearEntry := widget.NewEntry()
	dateLab := widget.NewLabel("Date")
	dateEntry := widget.NewEntry()
	infoTypeEntry := widget.NewSelect([]string{db.Creative.String(), db.Personal.String(), db.Professional.String()},
		func(val string) { selInfoType.SetText(val) })
	infoTypeLab := widget.NewLabel("InfoType")
	localeLab := widget.NewLabel("Locale")
	localeEntry := widget.NewEntry()
	imageLab := widget.NewLabel("Image")
	imageEntry := widget.NewButton("Select", func() {
		ChooseFile(w, selImageForInfo)
		// println(selImageForInfo)
	})
	albumsLab := widget.NewLabel("Albums")
	albumsEntry := widget.NewEntry()
	tracksLab := widget.NewLabel("Tracks")
	tracksEntry := widget.NewEntry()
	imageAlbumLab := widget.NewLabel("Image album")
	imageAlbumEntry := widget.NewEntry()
	saveInfoCard := widget.NewButton("Save Info Card", func() {})
	selected := container.NewVBox(widget.NewLabel(""), widget.NewLabel(""), widget.NewLabel(""), widget.NewLabel(""), selInfoType, widget.NewLabel(""), selImageForInfo)
	return container.NewBorder(nil, saveInfoCard, container.New(layout.NewGridLayoutWithColumns(2), titleLab, titleEntry,
		descrLab, descrEntry,
		yearLab, yearEntry,
		dateLab, dateEntry,
		infoTypeLab, infoTypeEntry,
		localeLab, localeEntry,
		imageLab, imageEntry,
		albumsLab, albumsEntry,
		tracksLab, tracksEntry,
		imageAlbumLab, imageAlbumEntry), selected)
}
