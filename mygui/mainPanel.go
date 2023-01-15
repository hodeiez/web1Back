package gui

import (
	"hodei/web1/db"
	"hodei/web1/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	tracksData  = [][]string{{"Key", "Name", "Description", "Date", "File"}}
	albumData   = []db.DbAlbumDTO{}
	albumsPanel = fyne.Container{}
)

func MainPanel(a fyne.App, w fyne.Window, panels [](func(wi fyne.Window) *fyne.Container)) *fyne.Container {
	toSaveTrack := widget.NewButton("SaveTrack", func() {
		tracksWin := a.NewWindow("tracks")
		tracksWin.SetContent(panels[0](w))
		tracksWin.Show()
	})

	toSaveAlbum := widget.NewButton("SaveAlbum", func() {
		albumsWin := a.NewWindow("Albums")
		albumsWin.SetContent(panels[1](w))
		albumsWin.Show()
	})
	toSaveInfoCard := widget.NewButton("SaveInfoCards", func() {
		infosWin := a.NewWindow("Infos")
		infosWin.SetContent(panels[2](w))
		infosWin.Show()

	})

	listTracks := widget.NewButton("ListTracks", func() {

		tracksData = setTracksInTable()

	})

	listAlbums := widget.NewButton("ListAlbums", func() {

		albumData = service.GetAllAlbums()
		albumsPanel = setGetAlbumsPanel(w)

	})
	listInfoCards := widget.NewButton("ListInfoCards", func() {})
	fetchAll := widget.NewButton("Fetch All", func() {})

	tracksTable := widget.NewTable(func() (int, int) { return MyTableLength(tracksData) }, MyCreateTable, MyUpdateTable)
	for i := 0; i < 5; i++ {
		tracksTable.SetColumnWidth(i, 100)
	}
	return container.New(layout.NewFormLayout(), container.New(layout.NewVBoxLayout(), toSaveTrack, toSaveAlbum, toSaveInfoCard, listTracks, listAlbums, listInfoCards, fetchAll), container.New(layout.NewGridLayout(2), tracksTable, container.NewPadded(&albumsPanel)))
}
func setTracksInTable() [][]string {
	fetched := service.GetAllTracks()
	returnData := make([][]string, len(fetched)+1)
	returnData[0] = tracksData[0]
	for i, t := range fetched {
		returnData[i+1] = []string{t.Key, t.Title, t.Description, t.Date, t.FileRef}
	}
	return returnData
}

func setGetAlbumsPanel(w fyne.Window) fyne.Container {

	hBoxes := make([]fyne.Container, len(albumData))
	vBox := *container.NewVBox()
	for i, a := range albumData {

		tracksList := widget.NewList(func() int { return len(a.Tracks) }, func() fyne.CanvasObject {
			return widget.NewButton("the size can be this long", func() {})
		}, func(i int, c fyne.CanvasObject) {
			c.(*widget.Button).SetText(a.Tracks[i].Title)
			c.(*widget.Button).OnTapped = func() {
				modal := TrackModal(w, a.Tracks[i])
				modal.Show()
			}
		})

		hBoxes[i] = *container.NewHBox(widget.NewLabel(a.Key), widget.NewLabel(a.Title), widget.NewLabel(a.Description), widget.NewLabel(a.Year), tracksList)
		vBox.Add(&hBoxes[i])
	}

	return vBox
}
