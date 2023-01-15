package gui

import (
	"hodei/web1/db"
	"hodei/web1/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var trackRefListForAlbum = binding.NewStringList()

func AlbumPanel(w fyne.Window) *fyne.Container {
	titleInp := widget.NewEntry()
	label1 := widget.NewLabel("Title")
	descrInp := widget.NewEntry()
	label2 := widget.NewLabel("Description")
	dateInp := widget.NewEntry()
	label3 := widget.NewLabel("Date")
	yearInp := widget.NewEntry()
	label6 := widget.NewLabel("Date")
	label4 := widget.NewLabel("Tracks")
	label5 := widget.NewLabel("")

	savedTracksList := widget.NewListWithData(trackRefListForAlbum,
		func() fyne.CanvasObject {

			return widget.NewButton("Response", func() {

				println("tapped")
			})
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			txt, _ := i.(binding.String).Get()
			o.(*widget.Button).OnTapped = func() {
				track, err := db.GetTrackByKey(o.(*widget.Button).Text)
				if err != nil {
					println(err)
				}
				mod := TrackModal(w, track)
				mod.Show()
			}
			o.(*widget.Button).SetText(txt)
		})
	data := binding.NewStringList()

	but := widget.NewButton("Save Album", func() {
		data.Set([]string{titleInp.Text, descrInp.Text, dateInp.Text, yearInp.Text})
		setData, _ := data.Get()
		trackList, _ := trackRefListForAlbum.Get()
		saved := saveAlbum(setData[0], setData[1], setData[2], setData[3], trackList)
		data.Set([]string{saved.Key, saved.Title, saved.Description, saved.Date})
		for _, track := range saved.Tracks {
			data.Append(track.Title)
		}
		// data.Set([]string{saved.Key, saved.Title, saved.Description, saved.Date, saved.FileRef})

	})

	addTrack := widget.NewEntry()
	adtrackRef := widget.NewButton("Add", func() {
		trackRefListForAlbum.Append(addTrack.Text)
	})

	result := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("Response")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			// o.(*widget.Label).Bind(i.(binding.String))
			txt, _ := i.(binding.String).Get()
			o.(*widget.Label).SetText(txt)

		})

	panel := container.New(layout.NewGridLayout(3), container.New(layout.NewFormLayout(), label1, titleInp, label2, descrInp, label3, dateInp, label6, yearInp, label4, container.NewHBox(addTrack, adtrackRef), label5, but), result, savedTracksList)
	panel.Resize(fyne.Size{500, 500})
	return panel
}

func saveAlbum(title string, descr string, dateInp string, yearInp string, tracks []string) db.DbAlbumDTO {
	return service.AddAlbum(db.DbAlbum{Title: title, Description: descr, Date: dateInp, Tracks: tracks, Year: yearInp})
}
