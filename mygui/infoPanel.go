// package gui

// import (
// 	"hodei/web1/db"
// 	"hodei/web1/service"
// 	"strings"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/layout"
// 	"fyne.io/fyne/v2/widget"
// )

// func InfoPanel(w fyne.Window) *fyne.Container {
// 	empty := widget.NewLabel("")
// 	selImageForInfo := widget.NewLabel("")
// 	selInfoType := widget.NewLabel("")
// 	selTracks := widget.NewLabel("")
// 	selAlbums := widget.NewLabel("")
// 	selImageAlbum := widget.NewLabel("")
// 	titleLab := widget.NewLabel("Title")
// 	titleEntry := widget.NewEntry()

// 	descrLab := widget.NewLabel("Description")
// 	descrEntry := widget.NewEntry()
// 	yearLab := widget.NewLabel("Year")
// 	yearEntry := widget.NewEntry()
// 	dateLab := widget.NewLabel("Date")
// 	dateEntry := widget.NewEntry()
// 	infoTypeEntry := widget.NewSelect([]string{db.Creative.String(), db.Personal.String(), db.Professional.String()},
// 		func(val string) { selInfoType.SetText(val) })
// 	infoTypeLab := widget.NewLabel("InfoType")
// 	localeLab := widget.NewLabel("Locale")
// 	localeEntry := widget.NewEntry()
// 	imageLab := widget.NewLabel("Image")
// 	imageEntry := widget.NewButton("Select", func() {
// 		ChooseFile(w, selImageForInfo)
// 	})
// 	albumsLab := widget.NewLabel("Albums")
// 	albumsEntry := widget.NewEntry()
// 	addAlbum := widget.NewButton("add 1", func() {
// 		addItemToLabel(albumsEntry.Text, selAlbums)
// 	})
// 	tracksLab := widget.NewLabel("Tracks")
// 	tracksEntry := widget.NewEntry()
// 	addTrack := widget.NewButton("add 1", func() {
// 		addItemToLabel(tracksEntry.Text, selTracks)
// 	})
// 	imageAlbumLab := widget.NewLabel("Image album")
// 	imageAlbumEntry := widget.NewEntry()
// 	addImageForAlbum := widget.NewButton("add 1", func() {
// 		addItemToLabel(imageAlbumEntry.Text, selImageAlbum)
// 	})
// 	saveInfoCard := widget.NewButton("Save Info Card", func() {
// 		service.AddInfoCard(db.DbInfo{Year: yearEntry.Text, Locale: localeEntry.Text, Title: titleEntry.Text, Description: descrEntry.Text, ImageRef: selImageForInfo.Text, AlbumsRef: toSlice(selAlbums.Text), TracksRef: toSlice(selTracks.Text), Date: dateEntry.Text, ImageAlbum: toSlice(selImageAlbum.Text), InfoType: selInfoType.Text})
// 	})
// 	selected := container.NewVBox(widget.NewLabel(""), widget.NewLabel(""), widget.NewLabel(""), widget.NewLabel(""), selInfoType, widget.NewLabel(""), selImageForInfo, selAlbums, selTracks, selImageAlbum)
// 	return container.NewBorder(nil, saveInfoCard, container.New(layout.NewGridLayoutWithColumns(3), titleLab, titleEntry, empty,
// 		descrLab, descrEntry, empty,
// 		yearLab, yearEntry, empty,
// 		dateLab, dateEntry, empty,
// 		infoTypeLab, infoTypeEntry, empty,
// 		localeLab, localeEntry, empty,
// 		imageLab, imageEntry, empty,
// 		albumsLab, albumsEntry, addAlbum,
// 		tracksLab, tracksEntry, addTrack,
// 		imageAlbumLab, imageAlbumEntry, addImageForAlbum), selected)
// }

// func addItemToLabel(item string, label *widget.Label) {
// 	if len(label.Text) > 0 {
// 		label.SetText(label.Text + ", " + item)
// 	} else {
// 		label.SetText(item)
// 	}
// }
// func toSlice(s string) []string {
// 	return strings.Split(s, ", ")
// }
