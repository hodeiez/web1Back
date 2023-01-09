package gui

import (
	"hodei/web1/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var tracksData = [][]string{{"Key", "Name", "Description", "Date", "File"}}

func MainPanel(a fyne.App, w fyne.Window, panel func(wi fyne.Window) *fyne.Container) *fyne.Container {
	toSaveTrack := widget.NewButton("SaveTrack", func() {
		tracksWin := a.NewWindow("tracks")
		tracksWin.SetContent(panel(w))
		tracksWin.Show()
	})

	toSaveAlbum := widget.NewButton("SaveAlbum", func() {})
	toSaveInfoCard := widget.NewButton("SaveInfoCards", func() {})
	listTracks := widget.NewButton("ListTracks", func() {

		tracksData = setTracksInTable()

	})
	listAlbums := widget.NewButton("ListAlbums", func() {})
	listInfoCards := widget.NewButton("ListInfoCards", func() {})
	fetchAll := widget.NewButton("Fetch All", func() {})

	tracksTable := widget.NewTable(func() (int, int) { return MyTableLength(tracksData) }, MyCreateTable, MyUpdateTable)

	return container.New(layout.NewFormLayout(), container.New(layout.NewVBoxLayout(), toSaveTrack, toSaveAlbum, toSaveInfoCard, listTracks, listAlbums, listInfoCards, fetchAll), container.New(layout.NewGridLayout(2), tracksTable))
}
func setTracksInTable() [][]string {
	fetched := service.GetAllTracks()
	returnData := make([][]string, len(fetched)+1)
	returnData[0] = tracksData[0]
	for i, t := range service.GetAllTracks() {
		returnData[i+1] = []string{t.Key, t.Title, t.Description, t.Date, t.FileRef}
	}
	return returnData
}
func albumMock() {

}
