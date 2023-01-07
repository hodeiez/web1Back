package gui

import (
	"fmt"
	"hodei/web1/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var tracksFetch = []string{"", "", "", "", ""}
var tracksData = [][]string{[]string{"Key", "Name", "Description", "Date", "File"},
	tracksFetch, tracksFetch}

func MainPanel(a fyne.App, w fyne.Window, panel func(wi fyne.Window) *fyne.Container) *fyne.Container {
	toSaveTrack := widget.NewButton("SaveTrack", func() {
		tracksWin := a.NewWindow("tracks")
		tracksWin.SetContent(panel(w))
		tracksWin.Show()
	})
	// fetchedTracks := binding.NewStringList()

	toSaveAlbum := widget.NewButton("SaveAlbum", func() {})
	toSaveInfoCard := widget.NewButton("SaveInfoCards", func() {})
	listTracks := widget.NewButton("ListTracks", func() { fmt.Println(service.GetAllTracks()) })
	listAlbums := widget.NewButton("ListAlbums", func() {})
	listInfoCards := widget.NewButton("ListInfoCards", func() {})
	fetchAll := widget.NewButton("Fetch All", func() {})
	tracksTable := widget.NewTable(func() (int, int) {
		return len(tracksData), len(tracksData[0])
	},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tracksData[i.Row][i.Col])
		})
	return container.New(layout.NewFormLayout(), container.New(layout.NewVBoxLayout(), toSaveTrack, toSaveAlbum, toSaveInfoCard, listTracks, listAlbums, listInfoCards, fetchAll), tracksTable)
}
